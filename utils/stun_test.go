package utils

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestStun(t *testing.T) {
	c, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Fatal(err)
	}
	// StunRequest("stun.l.google.com:19302", c)
	StunRequest("stun1.l.google.com:19302", c)
	StunRequest("stun2.l.google.com:19302", c)
	StunRequest("stun3.l.google.com:19302", c)
	go func() {
		for {
			buf := make([]byte, 1024)
			n, _, err := c.ReadFromUDP(buf)
			if err != nil {
				log.Fatal(err)
			}
			s, err := StunResolve(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
			log.Println(s)
		}
	}()
	time.Sleep(time.Second * 5)
	c.Close()
}
