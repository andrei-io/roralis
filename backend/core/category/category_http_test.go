package category_test

import (
	"backend/roralis/core/category"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

type categoryRepoMock struct {
	data          []category.Category
	notFoundError bool
}

func (r *categoryRepoMock) Get(id string) (*category.Category, error) {
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
func (r *categoryRepoMock) GetAll() ([]category.Category, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	return r.data, nil
}

func (r *categoryRepoMock) Create(c *category.Category) error {
	return repo.ErrNotImplementedYet

}
func (r *categoryRepoMock) Update(id string, c *category.Category) error {
	return repo.ErrNotImplementedYet
}
func (r *categoryRepoMock) Delete(id string) error {
	return repo.ErrNotImplementedYet
}

func TestCategoryController_ReadAll(t *testing.T) {
	t.Parallel()
	mockRepo := categoryRepoMock{
		notFoundError: false,
		data: []category.Category{
			{ID: 1, Text: "Categoria 1"},
			{ID: 2, Text: "Categoria 2"},
		},
	}
	c, w := rest.NewMockGinContext()
	categoryController := category.NewCategoryController(&mockRepo)
	categoryController.ReadAll(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	var responseSucces []category.Category
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if len(responseSucces) != len(mockRepo.data) {
		t.Errorf("Wanted number of items: %v, got %v", len(mockRepo.data), len(responseSucces))
	}

	mockRepo = categoryRepoMock{notFoundError: true}
	c, w = rest.NewMockGinContext()
	categoryController = category.NewCategoryController(&mockRepo)
	categoryController.ReadAll(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}

func TestCategoryController_ReadOne(t *testing.T) {
	t.Parallel()
	mockRepo := categoryRepoMock{
		notFoundError: false,
		data: []category.Category{
			{ID: 1, Text: "Categoria 1"},
			{ID: 2, Text: "Categoria 2"},
		},
	}
	c, w := rest.NewMockGinContext()
	categoryController := category.NewCategoryController(&mockRepo)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
	categoryController.ReadOne(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	var responseSucces category.Category
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if responseSucces != mockRepo.data[0] {
		t.Errorf("Wanted items: %v, got %v", mockRepo.data[0], responseSucces)
	}

	mockRepo = categoryRepoMock{notFoundError: true}
	c, w = rest.NewMockGinContext()
	categoryController = category.NewCategoryController(&mockRepo)
	categoryController.ReadAll(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}
