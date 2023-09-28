package randomstring

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	const ALPHANUMERIC = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = ALPHANUMERIC[rand.Intn(len(ALPHANUMERIC))]
	}
	return string(bytes)
}
