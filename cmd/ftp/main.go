package main

import (
	"net/url"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/yanggao40/goftp"
	"golang.org/x/net/proxy"
)

func main() {
	proxy.RegisterDialerType("http", goftp.NewHTTPProxy)
	proxy.RegisterDialerType("https", goftp.NewHTTPProxy)

	ftpProxySocks()
}

func local() {
	proxyURL, err := url.Parse("socks5://localhost:1080")
	if err != nil {
		panic(err)
	}

	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		panic(err)
	}

	conn, err := ftp.Dial("localhost:21",
		ftp.DialWithTimeout(30*time.Second),
		ftp.DialWithDisabledEPSV(true),
		ftp.DialWithDialFunc(dialer.Dial),
	)
	if err != nil {
		panic(err)
	}

	err = conn.Login("user", "123")
	if err != nil {
		panic(err)
	}

	err = conn.Quit()
	if err != nil {
		panic(err)
	}
}

func adeal() {
	proxyURL, err := url.Parse("http://proxy.adeal.ru:3128")
	if err != nil {
		panic(err)
	}

	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		panic(err)
	}

	conn, err := ftp.Dial("localhost:21",
		ftp.DialWithTimeout(30*time.Second),
		ftp.DialWithDisabledEPSV(true),
		ftp.DialWithDialFunc(dialer.Dial),
	)
	if err != nil {
		panic(err)
	}

	err = conn.Login("user", "123")
	if err != nil {
		panic(err)
	}

	err = conn.Quit()
	if err != nil {
		panic(err)
	}
}

func adealru() {
	proxyURL, err := url.Parse("http://proxyru.adeal.ru:3129")
	if err != nil {
		panic(err)
	}

	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		panic(err)
	}

	conn, err := ftp.Dial("localhost:21",
		ftp.DialWithTimeout(30*time.Second),
		ftp.DialWithDisabledEPSV(true),
		ftp.DialWithDialFunc(dialer.Dial),
	)
	if err != nil {
		panic(err)
	}

	err = conn.Login("user", "123")
	if err != nil {
		panic(err)
	}

	err = conn.Quit()
	if err != nil {
		panic(err)
	}
}

func ftpProxySocks() {
	proxyURL, err := url.Parse("socks5://ftp-proxy.adeal.ru:1080")
	if err != nil {
		panic(err)
	}

	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		panic(err)
	}

	conn, err := ftp.Dial("88.218.243.114:21",
		ftp.DialWithTimeout(30*time.Second),
		ftp.DialWithDisabledEPSV(true),
		ftp.DialWithDialFunc(dialer.Dial),
	)
	if err != nil {
		panic(err)
	}

	err = conn.Login("ftp_market_smm", "thi3VohVo2thahc6phe4")
	if err != nil {
		panic(err)
	}

	err = conn.Quit()
	if err != nil {
		panic(err)
	}
}
