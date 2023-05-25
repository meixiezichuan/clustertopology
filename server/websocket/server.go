package websocket

import (
	"crypto/tls"
	"github.com/gobwas/ws"
	"github.com/google/uuid"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const defaultMinConn = 10000
const defaultMaxConn = 2000000

type ServerConfig struct {
	Key  string
	Cert string
}

var tempDelay time.Duration // how long to sleep on accept failure

type Server struct {
	Addr            string
	tcpLn           *net.TCPListener
	tlsLn           net.Listener
	TLSEnable       bool
	eventHandlers   map[string]EventHandlers
	routeConns      map[string]AddrConns
	WSMessages      *MessageQueue
	messageFinished chan struct{}
	lock            sync.RWMutex
	elock           sync.RWMutex
	stop            bool
	Config          ServerConfig
	Authentication  Authentication
	handlerWG       sync.WaitGroup
	upgraderOpts    upgraderOptions
}

type Authentication func([]byte) error

type WSMessage struct {
	Message        []byte
	ServerConnName string
	Route          string
	OpCode         OpCode
}

type upgraderOptions struct {
	onheader        func(connId string, key, value []byte) (conninfo string, err error)
	onrequest       func(connId string, url []byte) error
	onprotocol      func(connId string, proto []byte) (string, bool)
	onbeforeupgrade func(connId string) (conninfo string, err error)
}

type EventHandlers map[EventType]EventHandler

func New(addr string, config ServerConfig) *Server {
	s := &Server{
		Addr:            addr,
		eventHandlers:   map[string]EventHandlers{},
		routeConns:      map[string]AddrConns{},
		WSMessages:      NewMessageQueue(),
		messageFinished: make(chan struct{}),
		Config:          config,
		upgraderOpts:    upgraderOptions{},
	}
	return s
}

func (s *Server) RegisterAuth(f Authentication) {
	s.Authentication = f
}

func (s *Server) RegisterOnheader(onheader func(uid string, key, value []byte) (conninfo string, err error)) {
	s.upgraderOpts.onheader = onheader
}

func (s *Server) RegisterOnRequest(onrequest func(uid string, url []byte) error) {
	s.upgraderOpts.onrequest = onrequest
}

func (s *Server) RegisterOnProtocol(onprotocol func(uid string, proto []byte) (string, bool)) {
	s.upgraderOpts.onprotocol = onprotocol
}

func (s *Server) RegisterOnBeforeUpgrade(onbeforeupgrade func(uid string) (conninfo string, err error)) {
	s.upgraderOpts.onbeforeupgrade = onbeforeupgrade
}

func (s *Server) Start() {
	var TLSConfig *tls.Config
	var err error
	if s.Config.Cert != "" || s.Config.Key != "" {
		TLSConfig = &tls.Config{
			NextProtos:   []string{"http/1.1"},
			Certificates: make([]tls.Certificate, 1),
		}
		TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(s.Config.Cert, s.Config.Key)
		if err != nil {
			wLog.Errorf("Load Certificate failed.")
			return
		}

		s.tlsLn, err = tls.Listen("tcp", s.Addr, TLSConfig)
		if err != nil {
			wLog.Errorf("Get tlsListener failed.")
			return
		}
		s.TLSEnable = true
	}
	if !s.TLSEnable {
		s.tcpLn, err = s.GetTCPListener()
		if err != nil {
			wLog.Errorf("Get tcpListener failed.")
			return
		}
	}

	for {
		var conn net.Conn
		// TLS connection disable tcp keepalive by default.
		// If not TLS, Use TCPConn's SetKeepAlive method to disable tcp keepalive.
		if s.stop {
			return
		}
		if s.TLSEnable {
			conn, err = s.tlsLn.Accept()
			if err != nil {
				wLog.Errorf("tcp accept error: %s", err)
				if CheckAndHandleTempError(err) {
					continue
				}
				break
			}
		} else {
			var tcpconn *net.TCPConn
			tcpconn, err = s.tcpLn.AcceptTCP()
			if err != nil {
				wLog.Errorf("tcp accept error: %s", err)
				if CheckAndHandleTempError(err) {
					continue
				}
				break
			}
			err1 := tcpconn.SetKeepAlive(false)
			if err1 != nil {
				wLog.Errorf("Disable tcp Keepalive failed.")
			}
			conn = tcpconn
		}
		tempDelay = 0
		if conn == nil {
			wLog.Errorf("get conn is nil")
			continue
		}
		s.handleTCPConn(conn)
	}
}

func (s *Server) GetTCPListener() (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", s.Addr)
	if err != nil {
		wLog.Errorf("Resolve address %s failed!", s.Addr)
		return nil, err
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		// handle error
		wLog.Errorf("Listen failed: %v", err)
		return nil, err
	}
	return ln, nil
}

func CheckAndHandleTempError(err error) bool {
	if ne, ok := err.(net.Error); ok && ne.Temporary() {
		if tempDelay == 0 {
			tempDelay = 5 * time.Millisecond
		} else {
			tempDelay *= 2
		}
		if max := 1 * time.Second; tempDelay > max {
			tempDelay = max
		}
		wLog.Infof("http: Accept error: %v; retrying in %v", err, tempDelay)
		time.Sleep(tempDelay)
		return true
	}
	return false
}

func (s *Server) GetMessageQueue() *MessageQueue {
	return s.WSMessages
}

func (s *Server) HandleMessages() {
	for {
		msg, shutdown := s.WSMessages.Get()
		if shutdown {
			s.messageFinished <- struct{}{}
			break
		}
		s.HandleMessage(msg)
	}
}

func (s *Server) HandleMessage(msg *WSMessage) {
	s.handlerWG.Add(1)
	defer func() {
		if r := recover(); r != nil {
			wLog.Errorf("Message: %v, handler panic: %v", msg, r)
		}
		s.handlerWG.Done()
	}()
	if msg == nil {
		wLog.Errorf("get msg nil")
		return
	}
	serverConn := s.GetConnByConnInfo(msg.ServerConnName)
	if serverConn != nil {
		eventContext := NewEventContext(serverConn.Conn, msg.Message, OpText)
		if serverConn.OnPush != nil {
			serverConn.OnPush(eventContext)
		}
	}
}

func (s *Server) HandleConnection(wsconn *ServerConn) {
	wsconnPath := wsconn.Path
	rAddr := wsconn.GetRemoteAddr()
	defer func() {
		mapKey := rAddr
		if wsconn.ConnInfo != "" {
			mapKey = wsconn.ConnInfo
		}
		s.lock.Lock()
		if conn, ok := s.routeConns[wsconnPath][mapKey]; ok {
			if conn.GetRemoteAddr() == rAddr {
				delete(s.routeConns[wsconnPath], mapKey)
			}
		}
		s.lock.Unlock()
	}()
	s.HandleWSConn(wsconn)
}

func (s *Server) RegisterEventHandler(path string, event EventType, handler EventHandler) {
	s.elock.Lock()
	defer s.elock.Unlock()
	if _, ok := s.eventHandlers[path]; ok {
		if _, ok := s.eventHandlers[path][event]; ok {
			wLog.Errorf("Path %s Event %s handler already exsit.", path, event)
			return
		}
		s.eventHandlers[path][event] = handler
	} else {
		handlers := EventHandlers{event: handler}
		s.eventHandlers[path] = handlers
	}
}

func (s *Server) HandleWSConn(wsconn *ServerConn) {
	// if no conhandler, get event handler
	defer func() {
		if e := recover(); e != nil {
			wLog.Errorf("recover in watch handler %v", e)
		}
	}()
	s.elock.RLock()
	eHandlers, ok := s.eventHandlers[wsconn.Path]
	s.elock.RUnlock()
	if ok {
		for event, eventHandler := range eHandlers {
			wsconn.AddEventListener(event, eventHandler)
		}
		go wsconn.KeepAlive()
		wsconn.Handle()
	} else {
		// if client use wrong path, there are no handlers for this connection, thus we should close it.
		_ = wsconn.CloseWithCode(CloseAbnormal, "Wrong url")
		wsconn.CloseLocalConn()
	}
}

func (s *Server) GetConnByConnInfo(connInfo string) *ServerConn {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, addrConns := range s.routeConns {
		serverConn, ok := addrConns[connInfo]
		if ok {
			return serverConn
		}
	}
	return nil
}

func (s *Server) getAllConnsSlices() [][]*ServerConn {
	s.lock.RLock()
	routeConns := make([][]*ServerConn, len(s.routeConns))
	i := 0
	for path := range s.routeConns {
		routeConns[i] = s.GetRouteConns(path)
		i++
	}
	s.lock.RUnlock()
	return routeConns
}

func (s *Server) GetAllConns() (wsConns []*ServerConn) {
	connsSlices := s.getAllConnsSlices()
	for _, conns := range connsSlices {
		wsConns = append(wsConns, conns...)
	}
	return wsConns
}

func (s *Server) GetRouteConns(path string) []*ServerConn {
	s.lock.RLock()
	defer s.lock.RUnlock()
	routeConns, ok := s.routeConns[path]
	if !ok {
		return []*ServerConn{}
	}
	conns := make([]*ServerConn, len(routeConns))
	i := 0
	for _, v := range routeConns {
		conns[i] = v
		i++
	}
	return conns
}

func (s *Server) Stop() {
	s.stop = true
	if s.TLSEnable {
		if s.tlsLn != nil {
			s.tlsLn.Close()
		}
	} else if s.tcpLn != nil {
		s.tcpLn.Close()
	}
	s.WSMessages.ShutDown()
	<-s.messageFinished
	s.handlerWG.Wait()
}

func (s *Server) handleTCPConn(conn net.Conn) {
	go s.upgrade(conn)
}

func (s *Server) upgrade(conn net.Conn) {
	header := ws.HandshakeHeaderHTTP(http.Header{
		"X-Go-Version": []string{runtime.Version()},
	})
	path := ""
	connInfo := ""
	connID := uuid.New().String()
	u := ws.Upgrader{
		OnBeforeUpgrade: func() (ws.HandshakeHeader, error) {
			var err error
			if s.upgraderOpts.onbeforeupgrade != nil {
				connInfo, err = s.upgraderOpts.onbeforeupgrade(connID)
			}
			return header, err
		},
		OnRequest: func(url []byte) error {
			path = string(url)
			if s.upgraderOpts.onrequest != nil {
				return s.upgraderOpts.onrequest(connID, url)
			}
			return nil
		},
		OnHeader: func(key, value []byte) error {
			if s.upgraderOpts.onheader != nil {
				connInfoHead, err := s.upgraderOpts.onheader(connID, key, value)
				if connInfoHead != "" {
					connInfo = connInfoHead
				}
				return err
			}
			return nil
		},
		Protocol: func(bytes []byte) bool {
			if s.upgraderOpts.onprotocol != nil {
				_, f := s.upgraderOpts.onprotocol(connID, bytes)
				return f
			}
			return true
		},
	}
	_, err := u.Upgrade(conn)
	if err != nil {
		conn.Close()
		wLog.Infof("upgrade error: %s, url: %s", err, path)
		return
	}
	wsconn := NewWSConn(conn)
	wsconn.Path = path
	mapKey := conn.RemoteAddr().String()
	if connInfo != "" {
		mapKey = connInfo
		wsconn.ConnInfo = connInfo
	}
	s.lock.Lock()
	if addrConns, ok := s.routeConns[wsconn.Path]; ok {
		// if connection with connInfo existed, close the old one, use the new
		if conn, ok := addrConns[mapKey]; ok {
			_ = conn.CloseWithCode(CloseForever, "User login on another connection.")
			conn.CloseLocalConn()
		}
		addrConns[mapKey] = wsconn
	} else {
		s.routeConns[wsconn.Path] = AddrConns{mapKey: wsconn}
	}
	s.lock.Unlock()
	s.HandleConnection(wsconn)
}
