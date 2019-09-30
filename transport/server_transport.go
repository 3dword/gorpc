package transport

import (
	"context"
	"github.com/diubrother/gorpc/codes"
	"net"
	"fmt"
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
			go handleConn(ctx , conn)
		}

	}
	return nil
}

func (s *serverTransport) ListenAndServeUdp(ctx context.Context, opts ...ServerTransportOption) error {

	return nil
}

func handleConn(ctx context.Context, conn net.Conn) error {

	return nil
}