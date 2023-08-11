package main

import "testing"

func TestServer(t *testing.T) {
	server("https://helper.moonchan.xyz/123", "localhost:9001")
}
