package util

import (
	"fmt"
	"math/rand"
	"strings"
)

func init() {

}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// generates a random name of 6 characters
func RandomOwner() string {
	return RandomString(6)
}

// generates a random account balance between 1 and 1000
func RandomMoney() int64 {
	return RandomInt(1, 1000)
}

// generates a random currency code
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomID() int64 {
	return RandomInt(0, 100)
}

// generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
