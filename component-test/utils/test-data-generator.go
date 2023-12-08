package utils

import (
	randCrypto "crypto/rand"
	"encoding/base64"
	"golang.org/x/exp/rand"
)

func GenerateRandomStringWithLen(length int) string {
	b := make([]byte, length)
	_, err := randCrypto.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func GenerateRandomString() string {
	return GenerateRandomStringWithLen(10)
}

func GenerateRandFloatsInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateRandFloats() float64 {
	return GenerateRandFloatsInRange(10, 100)
}
