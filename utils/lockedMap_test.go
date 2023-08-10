package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	m := NewLockedMap()
	m.Put("2312", 231)
	m.Put("2312", "231")
	m.Put("2313", "231")
	m.Put("2314", "231")
	r, _ := json.Marshal(m)
	fmt.Printf("%s", r)
}
