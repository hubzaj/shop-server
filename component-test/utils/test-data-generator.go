package utils

import (
	randCrypto "crypto/rand"
	"encoding/base64"
	"golang.org/x/exp/rand"
)

func GenerateRandStringWithLen(length int) string {
	b := make([]byte, length)
	_, err := randCrypto.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func GenerateRandString() string {
	return GenerateRandStringWithLen(10)
}

func GenerateRandFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateRandFloat() float64 {
	return GenerateRandFloatInRange(10, 100)
}
