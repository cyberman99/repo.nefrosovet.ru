package helpers

import (
	"math/rand"
	"time"
)

// RandomString generates random string of len n from chosen dictionary
func RandomString(n int, dict string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = dict[rand.Intn(len(dict))]
	}
	return string(b)
}
