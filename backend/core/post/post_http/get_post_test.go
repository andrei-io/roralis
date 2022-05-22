package posthttp_test

import (
	"backend/roralis/core/post"
	posthttp "backend/roralis/core/post/post_http"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

type postRepoMock struct {
	data          []post.Post
	notFoundError bool
}

func (r *postRepoMock) Get(id string) (*post.Post, error) {
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
func (r *postRepoMock) GetAll(offset int, limit int, newest bool) ([]post.Post, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	return r.data[offset:], nil
}

func (r *postRepoMock) GetByUserID(id string) ([]post.Post, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	var filtered []post.Post
	for _, v := range r.data {
		if strconv.Itoa(int(v.UserID)) == id {
			filtered = append(filtered, v)
		}
	}
	return filtered, nil
}

func (r *postRepoMock) Create(c *post.Post) error {
	return repo.ErrNotImplementedYet

}
func TestPostController_ReadAll(t *testing.T) {
	mockRepo := postRepoMock{
		notFoundError: false,
		data: []post.Post{
			{ID: 1, UserID: 1, Title: "First Post"},
			{ID: 2, UserID: 2, Title: "Second Post"},
		},
	}

	// Test GET all unfiltered
	c, w := rest.NewMockGinContext(nil)

	postController := posthttp.NewPostController(&mockRepo, "aaaa")

	postController.ReadAll(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	var responseSucces []post.Post
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if len(responseSucces) != len(mockRepo.data) {
		t.Errorf("Wanted number of items: %v, got %v", len(mockRepo.data), len(responseSucces))
	}

	// Test filtering(offset)
	c, w = rest.NewMockGinContext(&rest.TestHttpConfig{
		QueryParams: []rest.KV{
			{Key: "offset", Value: "1"},
		},
	})

	postController.ReadAll(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	err = json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if len(responseSucces) != len(mockRepo.data)-1 {
		t.Errorf("Wanted number of items: %v, got %v", len(mockRepo.data)-1, len(responseSucces))
	}

	c, w = rest.NewMockGinContext(&rest.TestHttpConfig{
		QueryParams: []rest.KV{
			{Key: "user_id", Value: "1"},
		},
	})

	// Test filtering(user_id)
	postController.ReadAll(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	err = json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if len(responseSucces) != 1 {
		t.Errorf("Wanted number of items: %v, got %v", 1, len(responseSucces))
	}

	// Test NotFound
	mockRepo.notFoundError = true
	c, w = rest.NewMockGinContext(nil)
	postController = posthttp.NewPostController(&mockRepo, "aaaa")
	postController.ReadAll(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}

}

func TestPostController_ReadOne(t *testing.T) {
	mockRepo := postRepoMock{
		notFoundError: false,
		data: []post.Post{
			{ID: 1, UserID: 1, Title: "First Post"},
			{ID: 2, UserID: 2, Title: "Second Post"},
		},
	}

	c, w := rest.NewMockGinContext(nil)

	postController := posthttp.NewPostController(&mockRepo, "aaaa")

	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	postController.ReadOne(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	var responseSucces post.Post
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if responseSucces != mockRepo.data[0] {
		t.Errorf("Wanted items: %v, got %v", mockRepo.data[0], responseSucces)
	}

	mockRepo.notFoundError = true
	c, w = rest.NewMockGinContext(nil)
	postController = posthttp.NewPostController(&mockRepo, "aaaa")
	postController.ReadOne(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}
