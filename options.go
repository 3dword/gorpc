package gorpc

type ServiceOptions struct {
	target string  // 监听地址，格式 ip://127.0.0.1:8080 , dns://www.google.com
}

type ServiceOption func(*ServiceOptions)

func WithTarget(target string) ServiceOption{
	return func(o *ServiceOptions) {
		o.target = target
	}
}

