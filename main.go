package main

import (
	"crypto/sha1"
	"log"
	"net"
	"time"

	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

func onError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	go server()
	client("")
}

func server() {

	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:12345")
	conn, _ := net.ListenUDP("udp", addr)
	go func() {
		for {
			time.Sleep(time.Second / 4)
			conn.WriteToUDP([]byte{}, addr)
		}
	}()

	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	listener, err := kcp.ServeConn(block, 10, 3, conn)
	onError(err)
	for {
		s, err := listener.AcceptKCP()
		onError(err)
		go handleEcho(s)
	}
}

// handleEcho send back everything it received
func handleEcho(conn *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// func client() {
// 	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
// 	block, _ := kcp.NewAESBlockCrypt(key)

// 	// wait for server to become ready
// 	time.Sleep(time.Second)

// 	conn, err := net.ListenUDP("udp", nil)
// 	onError(err)

// 	// dial to the echo server
// 	if sess, err := kcp.NewConn("127.0.0.1:12345", block, 10, 3, conn); err == nil {
// 		for {
// 			data := time.Now().String()
// 			buf := make([]byte, len(data))
// 			log.Println("sent:", data)
// 			if _, err := sess.Write([]byte(data)); err == nil {
// 				// read back the data
// 				if _, err := io.ReadFull(sess, buf); err == nil {
// 					log.Println("recv:", string(buf))
// 				} else {
// 					log.Fatal(err)
// 				}
// 			} else {
// 				log.Fatal(err)
// 			}
// 			time.Sleep(time.Second)
// 		}
// 	} else {
// 		log.Fatal(err)
// 	}
// }
