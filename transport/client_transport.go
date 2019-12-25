package transport

import (
	"context"
	"github.com/diubrother/gorpc/codes"
)

type clientTransport struct {
	opts *ClientTransportOptions
}

func (c *clientTransport) Send(ctx context.Context, req []byte, opts ...ClientTransportOption) error {
	if c.opts.NetworkType == "tcp" {
		return c.SendTcpReq(ctx, req)
	}

	if c.opts.NetworkType == "udp" {
		return c.SendUdpReq(ctx, req)
	}

	return codes.NetworkNotSupportedError
}

func (c *clientTransport) SendTcpReq(ctx context.Context, req []byte) error {

	// 从连接池里面获取一个连接


	return nil
}

func (c *clientTransport) SendUdpReq(ctx context.Context, req []byte) error {

	return nil
}