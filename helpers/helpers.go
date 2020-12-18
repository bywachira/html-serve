package helpers

import (
	"math/rand"
	"time"
)

// GeneratePortNumber handle generating port number
func GeneratePortNumber() int {
	rand.Seed(time.Now().UnixNano())

	min := 1000
	max := 99999

	port := rand.Intn(max-min+1) + min

	return port
}
