package auth_test

import (
	"backend/roralis/core/auth"
	"backend/roralis/core/jwt"
	"backend/roralis/core/user"
	"backend/roralis/shared/rest"
	"net/http"
	"testing"
)

func TestDecrypt(t *testing.T) {
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

	// Succesfully decode
	c, w := rest.NewMockGinContext(nil)
	c.Set("testing", &jwt.JWTClaims{
		Name:     testUser.Name,
		ID:       testUser.ID,
		Verified: testUser.Verified,
		Role:     testUser.Role,
	})
	authController.AboutMe(c)
	if w.Code != http.StatusOK {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	// Error on invalid shape
	c, w = rest.NewMockGinContext(nil)
	c.Set("testing", testUser)
	authController.AboutMe(c)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusInternalServerError, w.Code, w.Body.String())
	}

	// Error on missing claims
	c, w = rest.NewMockGinContext(nil)
	authController.AboutMe(c)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusInternalServerError, w.Code, w.Body.String())
	}

}
