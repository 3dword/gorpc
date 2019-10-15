package gorpc

import (
	"os"
	"os/signal"
	"syscall"
)

// gorpc Server, 一个 Server 可以拥有一个或者多个 service
type Server struct {
	services map[string]Service
}


func (s *Server) Register(serviceName string, service Service) {
	if serviceName == "" {
		return
	}
	s.services[serviceName] = service
}

func (s *Server) Serve() {

	for _, service := range s.services {
		go service.Serve()
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)
	<-ch

	s.Close()
}

func (s *Server) Close() {

}