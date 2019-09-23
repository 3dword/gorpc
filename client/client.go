package client

type Client interface {
	Dial(opt ...Options)
	Recv()
}

type Options interface{
	set(*Options)
}

type defaultClient struct {
	options *Options
}

func (c *defaultClient) Dial() {

}