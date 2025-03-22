package cacherpc

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/tymbaca/gorange/cmd/election/cache"
)

type Server struct {
	c *cache.Cache
}

func (c *Server) Prolong(req *ProlongReq, resp *ProlongResp) error {
	current, ok := c.c.Prolong(req.Key, req.TTL)
	resp.Current = current
	resp.Set = ok

	return nil
}

func (c *Server) Get(req *GetReq, resp *GetResp) error {
	current, _ := c.c.Get(req.Key)
	resp.Val = current

	return nil
}

func (c *Server) SetNX(req *SetNXReq, resp *SetNXResp) error {
	current, ok := c.c.SetNX(req.Key, req.Val, req.TTL)
	resp.Current = current
	resp.Set = ok

	return nil
}

func Serve(c *cache.Cache, addr string) error {
	server := &Server{c: c}
	rpc.Register(server)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return http.Serve(l, nil)
}
