package main

import "crypto/rand"

type Token struct {
	Username string
	Secret   []byte
}

func (token *Token) generate(username string) {
	token.Username = username
	token.Secret = make([]byte, 8)
	rand.Read(token.Secret)
}
