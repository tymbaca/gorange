package cacherpc

import (
	"net/rpc"
	"time"
)

type ProlongReq struct {
	Key string
	TTL time.Duration
}

type ProlongResp struct{}

type SetNXReq struct {
	Key, Val string
	TTL      time.Duration
}

type SetNXResp struct {
	Current string
	Set     bool
}

type Client struct {
	c *rpc.Client
}

func Connect(addr string) *Client {
	c, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		panic(err)
	}

	return &Client{c: c}
}

func (c *Client) Prolong(key string, ttl time.Duration) error {
	req := ProlongReq{Key: key, TTL: ttl}
	var resp ProlongResp
	err := c.c.Call("Server.Prolong", &req, &resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SetNX(key, val string, ttl time.Duration) (string, bool, error) {
	req := SetNXReq{Key: key, Val: val, TTL: ttl}
	var resp SetNXResp
	err := c.c.Call("Server.SetNX", &req, &resp)
	if err != nil {
		return "", false, err
	}

	return resp.Current, resp.Set, nil
}
