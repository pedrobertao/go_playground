package main

import (
	"crypto/rand"
	"math/big"
)

func randomGenerator(in string, size int) (string, error) {
	input := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomStr := make([]byte, size)
	encrypted := in + input

	for i := range randomStr {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(encrypted))))
		if err != nil {
			return "", err
		}
		randomStr[i] = encrypted[index.Int64()]
	}

	return string(randomStr), nil
}
