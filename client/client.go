package client

import "time"

// Client 定义了客户端通用接口
type Client interface {
	Invoke(rspbody interface{}, opts ...Option)
}

// Options 定义了客户端调用参数
type Options struct {
	// 调用地址
	target string
	// 超时时间
	timeout time.Duration

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

func (c *defaultClient) Dial() {

}