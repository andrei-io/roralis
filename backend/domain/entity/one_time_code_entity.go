package entity

import (
	"crypto/rand"
	"time"
)

// nolint: govet
type OneTimeCode struct {
	ID     uint64
	Code   string
	Active bool
	Expire time.Time
	UserID uint64
}

const otpChars = "1234567890"

// Generates a new verification code with a given length
func GenerateVerificationCode(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
