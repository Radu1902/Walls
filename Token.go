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
	token.Secret = string(randStr)
}
