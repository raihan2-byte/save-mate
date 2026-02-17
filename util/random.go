package util

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumberString(length int) string {
	rand.Seed(time.Now().UnixNano())

	result := ""

	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%d", rand.Intn(10))
	}

	return result
}
