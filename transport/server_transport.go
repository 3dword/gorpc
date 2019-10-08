package transport

import (
	"context"
	"github.com/diubrother/gorpc/codes"
	"net"
	"fmt"
	"time"
	"sync"
	"github.com/diubrother/gorpc/log"
)

type serverTransport struct {
	opts *ServerTransportOptions
}

func (s *serverTransport) ListenAndServe(ctx context.Context, opts ...ServerTransportOption) error {

	switch s.opts.Network {
		case "tcp","tcp4","tcp6":
			return s.ListenAndServeTcp(ctx, opts ...)
		case "udp","udp4", "udp6":
			return s.ListenAndServeUdp(ctx, opts ...)
		default:
			return codes.NewFrameworkError(102, "network not supported")
	}
}

func (s *serverTransport) ListenAndServeTcp(ctx context.Context, opts ...ServerTransportOption) error {
	for _, opt := range opts {
		opt(s.opts)
	}

	lis, err := net.Listen(s.opts.Network, s.opts.Address)
	if err != nil {
		return codes.NewFrameworkError(201, err.Error())
	}
	for {
		if conn , err := lis.Accept(); err != nil {
			return codes.NewFrameworkError(103,fmt.Sprintf("listener accept error, address : %s", s.opts.Address))
			go s.handleConn(ctx , conn)
		}

	}
	return nil
}

func (s *serverTransport) ListenAndServeUdp(ctx context.Context, opts ...ServerTransportOption) error {

	return nil
}

func (s *serverTransport) handleConn(ctx context.Context, rawConn net.Conn) error {
	rawConn.SetDeadline(time.Now().Add(s.opts.Timeout))
	tcpConn := newTcpConn(rawConn)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		s.read(ctx,tcpConn)
	}()
	go func() {
		defer wg.Done()
		s.handle(ctx,tcpConn)
	}()
	go func() {
		defer wg.Done()
		s.write(ctx,tcpConn)
	}()
	wg.Wait()

	return nil
}

func (s *serverTransport) read(ctx context.Context, conn *tcpConn) error {
	err := s.opts.Codec.Decode(conn.conn, msg)
	if err != nil {
		log.Error("read data from conn error, %v", err)
		return codes.ServerDecodeError
	}
	conn.in <- msg
	return nil
}

func (s *serverTransport) handle(ctx context.Context, conn *tcpConn) {
	req := <- conn.in

}

func (s *serverTransport) write(ctx context.Context, conn *tcpConn) {

}


type tcpConn struct {
	conn net.Conn
	in chan []byte
	out chan []byte
}

func newTcpConn(rawConn net.Conn) *tcpConn {
	return &tcpConn{
		conn : rawConn,
		in : make(chan []byte, 1),
		out : make(chan []byte, 1),
	}
}