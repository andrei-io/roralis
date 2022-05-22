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
	return r.data, nil
}

func (r *postRepoMock) Create(c *post.Post) error {
	return repo.ErrNotImplementedYet

}
func (r *postRepoMock) Update(id string, c *post.Post) error {
	return repo.ErrNotImplementedYet
}
func (r *postRepoMock) Delete(id string) error {
	return repo.ErrNotImplementedYet
}

func TestPostController_ReadAll(t *testing.T) {
	mockRepo := postRepoMock{
		notFoundError: false,
		data: []post.Post{
			{ID: 1, UserID: 1, Title: "First Post"},
			{ID: 1, UserID: 2, Title: "Second Post"},
		},
	}

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
}
