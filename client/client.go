package client

import "time"

// Client 定义了客户端通用接口
type Client interface {
	Invoke(rspbody interface{}, opts ...Option)
}

// Options 定义了客户端调用参数
type Options struct {
	// 调用地址
	Target string
	// 超时时间
	Timeout time.Duration

}

func WithTarget(target string) Option {
	return func(o *Options) {
		o.Target = target
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}



type Option func(*Options)


type defaultClient struct {
	options *Options
}

func (c *defaultClient) Dial() {

}