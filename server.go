package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/Hana-ame/tun-over-kcp/utils"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

func runConn(conn *net.UDPConn, url string) {
	for {
		time.Sleep(time.Second)
		if r, err := utils.Fetch("GET", url+"/list", nil, nil); err == nil {
			var l []string
			decoder := json.NewDecoder(r.Body)
			if err = decoder.Decode(&l); err == nil {
				for _, s := range l {
					addr, err := net.ResolveUDPAddr("udp", s)
					if err == nil {
						conn.WriteToUDP([]byte{}, addr)
					} else {
						fmt.Printf("%s\n", err)
					}
				}
			} else {
				fmt.Printf("%s\n", err)
			}
		} else {
			fmt.Printf("%s\n", err)
		}
	}
}

func server(url string, laddr string) {
	conn, err := net.ListenUDP("udp", nil)
	onError(err)
	addr, err := getStunIP(conn)
	for err != nil {
		fmt.Println(err)
		addr, err = getStunIP(conn)
	}
	fmt.Printf("addr: %s\n", addr)
	_, err = utils.Fetch("POST", url+"?addr="+addr, nil, nil)
	for err != nil {
		_, err = utils.Fetch("POST", url+"?addr="+addr, nil, nil)
	}
	go runConn(conn, url)

	// buf := make([]byte, 2048)
	// for {
	// 	n, addr, err := conn.ReadFromUDP(buf)
	// 	fmt.Println(n, err, addr)
	// }

	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	listener, err := kcp.ServeConn(block, 10, 3, conn)
	onError(err)
	for {
		s, err := listener.AcceptKCP()
		if err == nil {
			go handleKCP(s, laddr)
		} else {
			fmt.Printf("%s", err)
		}
	}
}

// handleEcho send back everything it received
func handleKCP(conn *kcp.UDPSession, laddr string) {
	fmt.Println("new client")

	proxy, err := net.Dial("tcp", laddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("proxy connected")
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}
