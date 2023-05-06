package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	var (
		letters string
		numeros string
		chars   string
	)

	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numeros = "0123456789"
	chars = "!@#$%^&*()+?><:{}[]"

	charsconcat := letters + numeros + chars

	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)
	for i := range password {
		password[i] = charsconcat[rand.Intn(len(charsconcat))]
	}

	return string(password)
}

func main() {
	password := generatePassword(20)
	fmt.Println(password)
}
