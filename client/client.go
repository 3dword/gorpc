package client

import (
	"context"
	"github.com/diubrother/gorpc/codes"
	"github.com/diubrother/gorpc/transport"
	"time"
)

// Client 定义了客户端通用接口
type Client interface {
	Invoke(ctx context.Context, req interface{}, rsp interface{}, opts ...Option) error
}

// Options 定义了客户端调用参数
type Options struct {
	// 调用地址
	target string
	// 超时时间
	timeout time.Duration

	Transport transport.ClientTransport
}

type Option func(*Options)

func WithTarget(target string) Option {
	return func(o *Options) {
		o.target = target
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

type defaultClient struct {
	options *Options
}

func (c *defaultClient) Invoke(ctx context.Context, req interface{}, rsp interface{}, opts ...Option) error {
	for _, opt := range opts {
		opt(c.options)
	}

	reqBytes, ok := req.([]byte)
	if !ok {
		return codes.ClientMsgError
	}

	if err := c.options.Transport.Send(ctx, reqBytes); err != nil {
		return codes.ClientNetworkError
	}

	return nil
}