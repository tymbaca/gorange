package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/url"

	"github.com/yanggao40/goftp"
	"golang.org/x/net/proxy"
)

// http://proxy.adeal.ru:3128
func main() {
	proxyURL, err := url.Parse("socks5://localhost:1080")
	if err != nil {
		panic(err)
	}

	proxy.RegisterDialerType("http", goftp.NewHTTPProxy)
	proxy.RegisterDialerType("https", goftp.NewHTTPProxy)

	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		panic(err)
	}

	//--------------------------------------------------------------------------------------------------

	// dialer, err = proxy.SOCKS5("tcp", "proxy.adeal.ru:1080", nil, proxy.Direct)
	// if err != nil {
	// 	panic(err)
	// }

	conn, err := dialer.Dial("tcp", "localhost:8001")
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	fmt.Println("sending")
	n, err := io.CopyN(conn, rand.Reader, 100)
	if err != nil {
		panic(err)
	}

	fmt.Printf("sent %d bytes \n", n)
}
