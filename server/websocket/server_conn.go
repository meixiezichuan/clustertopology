package websocket

import (
	"net"
	"time"
)

var KeepAliveInterval = 3 * time.Second
var LastPongTimeout = 3 * KeepAliveInterval

type ServerConn struct {
	*Conn
	Path     string
	IsClosed bool
}

type AddrConns map[string]*ServerConn

func NewWSConn(conn net.Conn) *ServerConn {
	c := NewServerConn(conn)
	wsc := &ServerConn{
		Conn:     c,
		IsClosed: false,
	}
	return wsc
}

func (s *ServerConn) KeepAlive() {
	ticker := time.NewTicker(KeepAliveInterval)
	defer ticker.Stop()
	failedTimes := 0
	firstPing := true
	var checkTime time.Time
	for range ticker.C {
		if s.IsClosed {
			break
		}

		// err checks whether directly connected peer is lost
		err := s.Write(OpPing, nil)
		if err != nil {
			failedTimes++
		} else {
			if failedTimes > 0 {
				failedTimes--
			}
			if firstPing {
				firstPing = false
				checkTime = time.Now()
			}
		}

		if failedTimes >= 3 {
			break
		}

		// LastPong checks whether end peer  is lost
		// if we never receive Pong, check how long since first ping sent
		if !s.LastPong.IsZero() {
			checkTime = s.LastPong
		}
		if time.Since(checkTime) > LastPongTimeout {
			break
		}
	}
	s.IsClosed = true
	s.CloseLocalConn()
}

func (s *ServerConn) Handle() {
	if s.OnOpen != nil {
		ctx := NewEventContext(s.Conn, nil, OpText)
		s.OnOpen(ctx)
	}
	var e error
	for {
		msg, _, err := s.read()
		if err != nil {
			s.CloseLocalConn()
			s.IsClosed = true
			e = err
			break
		}
		if s.OnMessage != nil {
			ctx := NewEventContext(s.Conn, msg, OpText)
			s.OnMessage(ctx)
		}
	}
	if IsCloseErr(e) {
		if s.OnClose != nil {
			ctx := s.GetErrorEventContext(e)
			s.OnClose(ctx)
		}
	} else {
		if s.OnError != nil {
			ctx := NewEventContext(s.Conn, []byte(e.Error()), OpText)
			s.OnError(ctx)
		}
	}
}
