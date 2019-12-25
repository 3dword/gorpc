package connpool

import (
	"context"
	"github.com/diubrother/gorpc/codes"
	"net"
	"sync"
)

type Pool interface {
	Get(ctx context.Context, network string, address string) (net.Conn, error)
}

type pool struct {
	opts *Options
	conns sync.Map
}

type channelPool struct {
	net.Conn
	initialCap int
	maxCap int
	Dial func(context.Context) (net.Conn, error)
	conns chan net.Conn
	mu sync.Mutex
}


func (p *pool) NewChannelPool(ctx context.Context, network string, address string) (*channelPool, error){
	c := &channelPool {
		initialCap: p.opts.initialCap,
		maxCap: p.opts.maxCap,
		Dial : func(ctx context.Context) (net.Conn, error) {
			return net.Dial(network, address)
		},
	}
	for i := 0; i < c.initialCap; i++ {
		conn , err := c.Dial(ctx);
		if err != nil {
			c.Close()
			return nil, codes.ConnectionPoolInitError
		}

		c.conns <- conn
	}

	return c, nil
}

func (c *channelPool) Get(ctx context.Context) (net.Conn, error) {
	if c.conns == nil {
		return nil, codes.ConnectionClosedError
	}
	select {
		case conn := <-c.conns :
			if conn == nil {
				return nil, codes.ConnectionClosedError
			}
			return c.wrapConn(conn), nil
		default:
			conn, err := c.Dial(ctx)
			if err != nil {
				return nil, codes.ClientNetworkError
			}
			return c.wrapConn(conn), nil
	}
}

func (c *channelPool) Close() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.Dial = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}
	close(conns)
	for conn := range conns {
		conn.Close()
	}
}

func (c *channelPool) Put(conn net.Conn) error {
	if conn == nil {
		return codes.ConnectionClosedError
	}
	if c.conns == nil {
		conn.Close()
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	select {
		case c.conns <- conn :
			return nil
	default:
		// 连接池满
		return conn.Close()
	}
}

func (p *pool) Get(ctx context.Context, network string, address string) (net.Conn, error) {

	if value ,ok := p.conns.Load(address); ok {
		if cp, ok := value.(*channelPool); ok {
			conn, err := cp.Get(ctx)
			return cp.wrapConn(conn), err
		}
	}

	cp, err := p.NewChannelPool(ctx, network, address)
	if err != nil {
		return nil, codes.ConnectionPoolInitError
	}

	p.conns.Store(address, cp)

	return cp.Get(ctx)
}

