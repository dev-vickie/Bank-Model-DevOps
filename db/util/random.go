// Package to generate random values for our unit tests
package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuioplkjhgfdsazxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder //string builder object
	k := len(alphabet)

	for i:=0;i<n;i++{
		c:= alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}


//sample use to gen ownern name
func RandomOwner()string {
	return RandomString(6)
}


//gen sample amount randomly
func RandomAmount() int64{
	return RandomInt(0,1000)
}

//choose random currency from the array
func RandomCurrency() string {
	currencies := []string{"EUR", "KES","USD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
