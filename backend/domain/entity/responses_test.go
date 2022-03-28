package entity_test

import (
	"backend/roralis/domain/entity"
	"testing"
)

func TestNewDuplicateEntityErrorResponse(t *testing.T) {
	tests := []struct {
		Field   string
		Correct string
	}{
		{
			Field:   "abc",
			Correct: "This value is already taken: abc",
		},
		{
			Field:   "users_email_key",
			Correct: "This value is already taken: users_email_key",
		},
		{
			Field:   "users_phone_number",
			Correct: "This value is already taken: users_phone_number",
		},
	}

	for _, tt := range tests {
		r := entity.NewDuplicateEntityErrorResponse(tt.Field)
		if r.Message != tt.Correct {
			t.Errorf("Failed.Got %q, Wanted %q", r.Message, tt.Correct)
		}
	}
}
