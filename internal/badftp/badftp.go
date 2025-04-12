package ftpclient

import (
	"io"
	"log"

	"github.com/jlaffaye/ftp"
)

type client struct {
	conn *ftp.ServerConn // держит коннекшн к хранилищу
}

func New(addr string) client {
	conn, err := ftp.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}

	return client{conn: conn}
}

func (c client) Read(path string, filename string) ([]byte, error) {
	resp, err := c.conn.Retr(path)
	if err != nil {
		return nil, err
	}
	resp.Close()

	return io.ReadAll(resp)
}
