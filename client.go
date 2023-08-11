package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Hana-ame/tun-over-kcp/utils"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

const N = 4

var (
	pool   = make(chan *net.UDPConn, N)
	addrs  = make(chan string, N)
	orders = make(chan struct{}, N)
	m      = make(map[string]*net.UDPConn)
	mu     sync.Mutex
	raddr  string
)

// works well
func newConn(url string) {
	defer func() {
		if err := recover(); err != nil {
			orders <- struct{}{}
		}
	}()
	conn, err := net.ListenUDP("udp", nil)
	onError(err)
	s, err := getStunIP(conn)
	onError(err)
	resp, err := utils.Fetch("GET", url+"?addr="+s, nil, nil)
	onError(err)
	if resp.StatusCode != 200 {
		onError(fmt.Errorf("%s", "not success"))
	}

	// got ip
	pool <- conn
	addrs <- s
	mu.Lock()
	m[s] = conn
	// fmt.Println("add", s)
	mu.Unlock()
}

// it should be good
func runPool(url string) {

	go func() {
		for i := 0; i < N; i++ {
			orders <- struct{}{}
		}
		for {
			<-orders
			newConn(url)
		}
	}()

	uaddr, _ := net.ResolveUDPAddr("udp", "8.8.8.8:53")
	for {
		time.Sleep(time.Second)
		mu.Lock()
		// fmt.Println(m)
		for key := range m {
			m[key].WriteToUDP([]byte{}, uaddr)
			resp, err := utils.Fetch("GET", url+"?addr="+key, nil, nil)
			// fmt.Println("?raddr", raddr, err)
			if err == nil {
				decoder := json.NewDecoder(resp.Body)
				_ = decoder.Decode(&raddr)
			}
		}
		mu.Unlock()
	}
}

func pick() *net.UDPConn {
	addr := <-addrs
	conn := <-pool
	mu.Lock()
	delete(m, addr)
	fmt.Printf("picked @ %s\n", addr)
	mu.Unlock()
	orders <- struct{}{}
	return conn
}

func clientDial() (*kcp.UDPSession, error) {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	conn := pick()
	fmt.Printf("dial %s -> %s\n", conn.LocalAddr().String(), raddr)
	// io.ReadAll(conn)
	// dial to the echo server
	return kcp.NewConn(raddr, block, 10, 3, conn)
}

func client(url string, laddr string) {
	go runPool(url)
	ln, err := net.Listen("tcp", laddr)
	onError(err)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	proxy, err := clientDial()
	onError(err)

	fmt.Printf("%s -> %s\n", conn.LocalAddr().String(), proxy.LocalAddr().String())
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}
