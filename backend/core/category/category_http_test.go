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

type regionRepoMock struct {
	data          []category.Category
	notFoundError bool
}

func (r *regionRepoMock) Get(id string) (*category.Category, error) {
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
func (r *regionRepoMock) GetAll() ([]category.Category, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	return r.data, nil
}

func (r *regionRepoMock) Create(c *category.Category) error {
	return repo.ErrNotImplementedYet

}
func (r *regionRepoMock) Update(id string, c *category.Category) error {
	return repo.ErrNotImplementedYet
}
func (r *regionRepoMock) Delete(id string) error {
	return repo.ErrNotImplementedYet
}

func TestCategoryController_ReadAll(t *testing.T) {

	mockRepo := regionRepoMock{
		notFoundError: false,
		data: []category.Category{
			{ID: 1, Text: "Categoria 1"},
			{ID: 2, Text: "Categoria 2"},
		},
	}
	c, w := rest.NewMockGinContext()
	categoryController := category.NewCategoryController(&mockRepo)

	// Succesfully get all categories
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

	// Artificially simulate a db failure(record not found)
	mockRepo = regionRepoMock{notFoundError: true}
	c, w = rest.NewMockGinContext()
	categoryController = category.NewCategoryController(&mockRepo)
	categoryController.ReadAll(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}

func TestCategoryController_ReadOne(t *testing.T) {

	mockRepo := regionRepoMock{
		notFoundError: false,
		data: []category.Category{
			{ID: 1, Text: "Categoria 1"},
			{ID: 2, Text: "Categoria 2"},
		},
	}

	c, w := rest.NewMockGinContext()
	categoryController := category.NewCategoryController(&mockRepo)

	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Sucesfully get controller by id
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

	c, w = rest.NewMockGinContext()

	// Error on invalid id
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "3"})
	categoryController.ReadOne(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}