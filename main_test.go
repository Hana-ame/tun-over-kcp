package main

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestStun(t *testing.T) {
	c, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Fatal(err)
	}
	s, err := getStunIP(c)
	fmt.Println(s, err)
}
