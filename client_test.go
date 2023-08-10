package main

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	go runPool("")
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
