package auth_test

import (
	"backend/roralis/core/auth"
	"backend/roralis/core/jwt"
	"backend/roralis/core/user"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type userRepoMock struct {
	data          []user.User
	notFoundError bool
}

func (r *userRepoMock) Get(id string) (*user.User, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	for _, v := range r.data {
		if strconv.Itoa(int(v.ID)) == id {
			return &v, nil
		}
	}
	return nil, repo.ErrRecordNotFound
}
func (r *userRepoMock) GetAll() ([]user.User, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	return r.data, nil
}

func (r *userRepoMock) GetByEmail(email string) (*user.User, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	for _, v := range r.data {
		if v.Email == email {
			return &v, nil
		}
	}
	return nil, repo.ErrRecordNotFound
}

func (r *userRepoMock) Create(c *user.User) error {
	return repo.ErrNotImplementedYet

}

type jwtServiceMock struct {
}

func (j *jwtServiceMock) NewJWT(c *jwt.JWTClaims) (string, error) {
	return "", nil
}
func (j *jwtServiceMock) VerifyJWT(token *string) (*jwt.JWTClaims, error) {
	return nil, nil
}

func TestSignIn(t *testing.T) {
	testUser := user.User{
		ID:       1,
		Name:     "FirstUser",
		Email:    "first@example.com",
		Password: "DefinetlyHashedPassword",
		Verified: true,
		Role:     5,
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Error on hasing password: %+v", err)
	}

	mockRepo := userRepoMock{
		notFoundError: false,
		data: []user.User{
			{
				ID:       testUser.ID,
				Name:     testUser.Name,
				Email:    testUser.Email,
				Password: string(hashedPassword),
				Verified: testUser.Verified,
				Role:     testUser.Role,
			},
		},
	}
	mockJWT := jwtServiceMock{}
	authController := auth.NewAuthController(&mockRepo, &mockJWT, "testing")

	// Succesful Sign In
	body, err := json.Marshal(auth.SignInRequest{Email: testUser.Email, Password: testUser.Password})
	if err != nil {
		t.Errorf("Error on marshalling json: %+v", err)
	}
	c, w := rest.NewMockGinContext(&rest.TestHttpConfig{Body: body})
	authController.SignIn(c)
	if w.Code != http.StatusOK {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusOK, w.Code, w.Body.String())
	}

	// Error on wrong password
	body, err = json.Marshal(auth.SignInRequest{Email: testUser.Email, Password: testUser.Password + "0"})
	if err != nil {
		t.Errorf("Error on marshalling json: %+v", err)
	}
	c, w = rest.NewMockGinContext(&rest.TestHttpConfig{Body: body})
	authController.SignIn(c)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusUnauthorized, w.Code, w.Body.String())
	}

}
