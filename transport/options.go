package transport

import "time"

type ServerTransportOptions struct{
	Address string // 地址，格式例如 ip://127.0.0.1：8080
	Network string  // 网络类型
	Timeout time.Duration  // 传输层请求超时时间，默认为 2 min
	MsgReader
}

type ServerTransportOption func(*ServerTransportOptions)

type ClientTransportOptions struct {
	Target string
}

type ClientTransportOption func(*ClientTransportOptions)
