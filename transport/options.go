package transport

type ServerTransportOptions struct{
	Address string // 地址，格式例如 ip://127.0.0.1：8080
	Network string  // 网络类型
}

type ServerTransportOption func(*ServerTransportOptions)

type ClientTransportOptions struct {
	Target string
}

type ClientTransportOption func(*ClientTransportOptions)
