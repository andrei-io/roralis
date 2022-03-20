package entity_test

import (
	"backend/roralis/domain/entity"
	"testing"
)

func TestGenerateOTP(t *testing.T) {
	for i := 0; i < 6; i++ {
		code, err := entity.GenerateVerificationCode(6)
		if err != nil {
			t.Errorf("Failed test with error: %s\n", err.Error())
		}
		if len(code) != 6 {
			t.Errorf("Failed test: code is not appropiate length, wanted 6,got %s", err.Error())
		}
	}
}
