package gorpc

import (
	"context"
	"github.com/diubrother/gorpc/log"
)

// Service 定义了某个具体服务的通用实现接口
type Service interface {
	Register()
	Serve()
	Close()
}


type service struct{
	serviceName string
	handlers map[string]Handler
	opts ServiceOptions
}

type Handler func(context.Context)


func (s *service) Register(handlerName string, handler Handler) {
	s.handlers[handlerName] = handler
}

func (s *service) Serve() {
	if s.opts.target == "" {
		log.Error("server listen address is empty")
		return
	}

}


func (s *service) Close() {

}

