package main

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client("https://helper.moonchan.xyz/123", ":8000")
}

func TestClientPool(t *testing.T) {
	go runPool("https://helper.moonchan.xyz/123")
	time.Sleep(time.Second)
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	time.Sleep(time.Second)
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	time.Sleep(time.Second)
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	fmt.Println(pick())
	time.Sleep(time.Second * 30)

}
