package gorpc

// gorpc Server, 一个 Server 可以拥有一个或者多个 service
type Server struct {
	services map[string]Service
}
