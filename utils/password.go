package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomPassword(length int) string {
	randomString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	password := ""
	for index := 0; index < length; index++ {
		randomIndex := rand.Intn(len(randomString))
		password += string(randomString[randomIndex])
	}

	return password
}
