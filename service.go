package gorpc

import "context"

// Service 定义了某个具体服务的通用实现接口
type Service interface {
	Register()
	Serve()
	Close()
}


type service struct{
	serviceName string
	handlers map[string]Handler
}

type Handler func(context.Context)


func (s *service) Register(handlerName string, handler Handler) {
	s.handlers[handlerName] = handler
}

func (s *service) Serve() {

}


func (s *service) Close() {

}

