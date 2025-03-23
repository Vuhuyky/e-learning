package random

import (
	"fmt"
	"math/rand"
)

func RandomCode(length int) string {
	b := ""
	for i := 0; i < length; i++ {
		b += fmt.Sprintf("%d", rand.Intn(9))
	}
	return b
}
