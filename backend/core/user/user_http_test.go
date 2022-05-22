package user_test

import (
	"backend/roralis/core/user"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
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
	return nil, repo.ErrNotImplementedYet
}

func (r *userRepoMock) Create(c *user.User) error {
	return repo.ErrNotImplementedYet

}

func TestUserController_ReadOne(t *testing.T) {
	mockRepo := userRepoMock{
		notFoundError: false,
		data: []user.User{
			{
				ID:       1,
				Name:     "FirstUser",
				Email:    "first@example.com",
				Password: "DefinetlyHashedPassword",
				Verified: true,
				Role:     5,
			},
			{
				ID:       2,
				Name:     "SecondUser",
				Email:    "second@example.com",
				Password: "DefinetlyHashedPassword2",
				Verified: true,
				Role:     5,
			},
		},
	}

	// Succesfully read a user by id
	c, w := rest.NewMockGinContext(nil)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	userController := user.NewUserController(&mockRepo)
	userController.ReadOne(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	var responseSucces user.User
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if responseSucces.Password == mockRepo.data[0].Password || responseSucces.Email == mockRepo.data[0].Email {
		t.Errorf("Wanted items: %v, got %v", mockRepo.data[0], responseSucces)
	}

	c, w = rest.NewMockGinContext(nil)

	// Error on invalid id
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "3"})
	userController.ReadOne(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}
