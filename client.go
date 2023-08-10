package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Hana-ame/tun-over-kcp/utils"
)

const N = 4

var (
	pool   = make(chan *net.UDPConn, N)
	addrs  = make(chan string, N)
	orders = make(chan struct{}, N)
	m      = make(map[string]*net.UDPConn)
	mu     sync.Mutex
)

func newConn() {
	defer func() {
		if err := recover(); err != nil {
			orders <- struct{}{}
		}
	}()
	buf := make([]byte, 1024)
	conn, err := net.ListenUDP("udp", nil)
	onError(err)
	ch := make(chan string)
	go func() {
		utils.StunRequest("stun1.l.google.com:19302", conn)
		n, _, err := conn.ReadFrom(buf)
		onError(err)
		s, err := utils.StunResolve(buf[:n])
		onError(err)
		ch <- s
	}()
	select {
	case <-time.After(3 * time.Second):
		defer conn.Close()
		onError(fmt.Errorf("UDP read timeout"))
	case s := <-ch:
		pool <- conn
		addrs <- s
		mu.Lock()
		m[s] = conn
		fmt.Println("add", s)
		mu.Unlock()
	}

}

func runPool(raddr string) {
	for i := 0; i < N; i++ {
		orders <- struct{}{}
	}
	go func() {
		for {
			<-orders
			newConn()
		}
	}()

	uaddr, _ := net.ResolveUDPAddr("udp", "8.8.8.8:53")
	for {
		time.Sleep(time.Second)
		mu.Lock()
		// fmt.Println(m)
		for key := range m {
			m[key].WriteToUDP([]byte{}, uaddr)
		}
		mu.Unlock()
	}
}

func pick() *net.UDPConn {
	addr := <-addrs
	conn := <-pool
	mu.Lock()
	delete(m, addr)
	fmt.Println("delete", addr)
	mu.Unlock()
	orders <- struct{}{}
	return conn
}

func client(raddr string) {
	go runPool(raddr)

}
