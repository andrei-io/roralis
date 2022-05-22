package region_test

import (
	"backend/roralis/core/region"
	"backend/roralis/shared/repo"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

type regionRepoMock struct {
	data          []region.Region
	notFoundError bool
}

func (r *regionRepoMock) Get(id string) (*region.Region, error) {
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
func (r *regionRepoMock) GetAll() ([]region.Region, error) {
	if r.notFoundError {
		return nil, repo.ErrRecordNotFound
	}
	return r.data, nil
}

func (r *regionRepoMock) Create(c *region.Region) error {
	return repo.ErrNotImplementedYet

}
func (r *regionRepoMock) Update(id string, c *region.Region) error {
	return repo.ErrNotImplementedYet
}
func (r *regionRepoMock) Delete(id string) error {
	return repo.ErrNotImplementedYet
}

func TestRegionController_ReadAll(t *testing.T) {

	mockRepo := regionRepoMock{
		notFoundError: false,
		data: []region.Region{
			{ID: 1, Text: "Regiunea 1"},
			{ID: 2, Text: "Regiunea 2"},
		},
	}
	c, w := rest.NewMockGinContext(nil)
	regionController := region.NewRegionController(&mockRepo)

	// Succesfully get all regions
	regionController.ReadAll(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}

	var responseSucces []region.Region
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if len(responseSucces) != len(mockRepo.data) {
		t.Errorf("Wanted number of items: %v, got %v", len(mockRepo.data), len(responseSucces))
	}

	// Artificially simulate a db failure(record not found)
	mockRepo = regionRepoMock{notFoundError: true}
	c, w = rest.NewMockGinContext(nil)
	regionController = region.NewRegionController(&mockRepo)

	regionController.ReadAll(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}

func TestRegionController_ReadOne(t *testing.T) {
	mockRepo := regionRepoMock{
		notFoundError: false,
		data: []region.Region{
			{ID: 1, Text: "Regiunea 1"},
			{ID: 2, Text: "Regiunea 2"},
		},
	}
	c, w := rest.NewMockGinContext(nil)
	regionController := region.NewRegionController(&mockRepo)

	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Succesfully get region by id
	regionController.ReadOne(c)
	if w.Code != http.StatusOK {
		t.Errorf("Wanted return code: %v, got %v", http.StatusOK, w.Code)
	}
	var responseSucces region.Region
	err := json.Unmarshal(w.Body.Bytes(), &responseSucces)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}
	if responseSucces != mockRepo.data[0] {
		t.Errorf("Wanted items: %v, got %v", mockRepo.data[0], responseSucces)
	}

	c, w = rest.NewMockGinContext(nil)

	// Error on invalid id
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "3"})
	regionController.ReadOne(c)
	if w.Code != http.StatusNotFound {
		t.Errorf("Wanted return code: %v, got %v", http.StatusNotFound, w.Code)
	}
}
