package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Hana-ame/tun-over-kcp/utils"
)

func onError(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	asServer bool
	asClient bool
	laddr    string
	url      string
)

func main() {
	flag.StringVar(&laddr, "laddr", "127.0.0.1:22", "tcp listen port")
	flag.StringVar(&url, "url", "https://helper.moonchan.xyz/123", "helper endpoint")
	flag.BoolVar(&asClient, "client", false, "as client")
	flag.BoolVar(&asServer, "server", false, "as server")

	flag.Parse()

	if asClient {
		client(url, laddr)
	} else if asServer {
		server(url, laddr)
	} else {
		fmt.Println("please set -server or -client")
	}
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}

func getStunIP(conn *net.UDPConn) (addr string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()

	// return fmt.Sprintf("127.0.0.1:%d", conn.LocalAddr().(*net.UDPAddr).Port), nil

	ch := make(chan string)
	go func() {
		buf := make([]byte, 1024)
		utils.StunRequest("stun1.l.google.com:19302", conn)
		n, _, err := conn.ReadFrom(buf)
		onError(err)
		s, err := utils.StunResolve(buf[:n])
		onError(err)
		ch <- s
	}()
	select {
	case <-time.After(3 * time.Second):
		onError(fmt.Errorf("UDP read timeout"))
	case s := <-ch:
		addr = s
	}
	return
}
