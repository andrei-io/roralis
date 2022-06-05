package auth_test

import (
	"backend/roralis/core/auth"
	"backend/roralis/core/user"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {
	testUser := user.User{
		ID:       1,
		Name:     "FirstUser",
		Email:    "first@example.com",
		Password: "DefinetlyHashedPassword",
		Verified: true,
		Role:     5,
	}
	mockRepo := userRepoMock{
		notFoundError: false,
		data:          []user.User{},
	}
	mockJWT := jwtServiceMock{}
	authController := auth.NewAuthController(&mockRepo, &mockJWT, "testing")

	// Succesful Sign Up
	body, err := json.Marshal(testUser)
	if err != nil {
		t.Errorf("Error on marshalling json: %+v", err)
	}
	c, w := rest.NewMockGinContext(&rest.TestHttpConfig{Body: body})
	authController.SignUp(c)
	if w.Code != http.StatusOK {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	// Error on duplicate email
	body, err = json.Marshal(testUser)
	if err != nil {
		t.Errorf("Error on marshalling json: %+v", err)
	}
	c, w = rest.NewMockGinContext(&rest.TestHttpConfig{Body: body})
	authController.SignUp(c)
	if w.Code != http.StatusConflict {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusConflict, w.Code, w.Body.String())
	}
}
