package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateReference() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	random := fmt.Sprintf("%04d", rand.Intn(10000))
	reference := fmt.Sprintf("EASYCOMMERCE-%d-%s", timestamp, random)
	return reference[:15]
}
