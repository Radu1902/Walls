package main

import "crypto/rand"

type Token struct {
	Username string
	Secret   string
}

func (token *Token) generate(username string) {
	token.Username = username
	randStr := make([]byte, 8)
	rand.Read(randStr)
	for index := range randStr {
		randStr[index] = (randStr[index] % 25) + 65
	}
	token.Secret = string(randStr)
}
