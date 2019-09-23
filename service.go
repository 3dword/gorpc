package gorpc

// Service 定义了某个具体服务的通用实现接口
type Service interface {
	Register()
	Serve()
	Close()
}
