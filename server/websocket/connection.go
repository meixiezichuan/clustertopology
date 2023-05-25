package websocket

import (
	"encoding/binary"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"io"
	"net"
	"strings"
	"time"
)

type Conn struct {
	netConn   net.Conn
	Reader    *wsutil.Reader
	LastPong  time.Time
	creatTime time.Time
	State     ws.State
	OnOpen    EventHandler
	OnClose   EventHandler
	OnMessage EventHandler
	OnError   EventHandler
	OnPush    EventHandler
	ConnInfo  string
}

type EventContext struct {
	Msg  []byte
	Code OpCode
	Conn *Conn
}

type EventHandler func(e *EventContext)

type EventType string
type CloseCode ws.StatusCode
type OpCode ws.OpCode

const (
	EventOpen    EventType = "open"
	EventClose   EventType = "close"
	EventMessage EventType = "message"
	EventPush    EventType = "push"
	EventError   EventType = "error"
)

const (
	CloseNormal   = CloseCode(ws.StatusNormalClosure)
	CloseForever  = CloseCode(ws.StatusGoingAway)
	CloseAbnormal = CloseCode(ws.StatusAbnormalClosure)
)

const (
	OpContinuation = OpCode(ws.OpContinuation)
	OpText         = OpCode(ws.OpText)
	OpBinary       = OpCode(ws.OpBinary)
	OpClose        = OpCode(ws.OpClose)
	OpPing         = OpCode(ws.OpPing)
	OpPong         = OpCode(ws.OpPong)
)

func NewWSReader(conn io.ReadWriter, s ws.State) *wsutil.Reader {
	ch := wsutil.ControlFrameHandler(conn, s)
	r := &wsutil.Reader{
		Source:         conn,
		State:          s,
		CheckUTF8:      true,
		OnIntermediate: ch,
	}
	return r
}

func NewServerConn(conn net.Conn) *Conn {
	return NewConn(conn, ws.StateServerSide)
}

func NewClientConn(conn net.Conn) *Conn {
	return NewConn(conn, ws.StateClientSide)
}

func NewConn(conn net.Conn, s ws.State) *Conn {
	reader := NewWSReader(conn, s)
	c := &Conn{
		creatTime: time.Now(),
		netConn:   conn,
		Reader:    reader,
		State:     s,
	}
	return c
}

func (c *Conn) customHandle(hdr ws.Header, err error) error {
	// if receive close frame, whatever the err is, return ClosedError to invoke OnClose
	if hdr.OpCode == ws.OpClose {
		if !IsCloseErr(err) {
			err = wsutil.ClosedError{
				Code: ws.StatusNoStatusRcvd,
			}
		}
		return err
	}
	// if get write error when reading, continue to read
	if err != nil {
		isWriteErr := strings.Contains(err.Error(), "write: broken pipe")
		if !isWriteErr {
			return err
		}
	}

	// if get pong, record pong time
	if hdr.OpCode == ws.OpPong {
		c.LastPong = time.Now()
	}
	return nil
}

func (c *Conn) read() ([]byte, byte, error) {
	return readData(c.Reader, c.customHandle)
}

func (c *Conn) Write(op OpCode, msg []byte) error {
	return writeData(c.netConn, ws.OpCode(op), msg, c.State)
}

func (c *Conn) WriteText(msg []byte) error {
	return writeData(c.netConn, ws.OpText, msg, c.State)
}

func (c *Conn) WriteBinary(msg []byte) error {
	return writeData(c.netConn, ws.OpBinary, msg, c.State)
}

func (c *Conn) Close(reason string) error {
	return closeNormal(c.netConn, reason, c.State)
}

func (c *Conn) CloseWithCode(code CloseCode, reason string) error {
	return writeClose(c.netConn, ws.StatusCode(code), reason, c.State)
}

func (c *Conn) CloseLocalConn() {
	c.netConn.Close()
}

func (c *Conn) GetRemoteAddr() string {
	return c.netConn.RemoteAddr().String()
}

func (c *Conn) GetLocalAddr() string {
	return c.netConn.LocalAddr().String()
}

func (c *Conn) GetCreateTime() time.Time {
	return c.creatTime
}

func (c *Conn) AddEventListener(event EventType, eventhandler EventHandler) {
	switch event {
	case EventOpen:
		c.OnOpen = eventhandler
	case EventClose:
		c.OnClose = eventhandler
	case EventMessage:
		c.OnMessage = eventhandler
	case EventPush:
		c.OnPush = eventhandler
	case EventError:
		c.OnError = eventhandler
	}
}

func (c *Conn) GetErrorEventContext(err error) *EventContext {
	p := make([]byte, 2)
	binary.BigEndian.PutUint16(p, uint16(err.(wsutil.ClosedError).Code))
	return NewEventContext(c, p, OpClose)
}

func NewEventContext(c *Conn, msg []byte, code OpCode) *EventContext {
	ctx := EventContext{
		Msg:  msg,
		Conn: c,
		Code: code,
	}
	return &ctx
}

func (e *EventContext) GetCloseCode() uint16 {
	if len(e.Msg) != 2 {
		return 0
	}
	return binary.BigEndian.Uint16(e.Msg)
}
