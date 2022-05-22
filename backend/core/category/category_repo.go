package category

import (
	"backend/roralis/shared/repo"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	GetAll() (c []Category, err error)
	Get(id string) (c *Category, err error)
}

type categoryRepo struct {
	db *gorm.DB
}

// Check interface at compile time
var _ CategoryRepo = (*categoryRepo)(nil)

// Constructr funcion
func NewCategoryRepo(db *gorm.DB) *categoryRepo {
	return &categoryRepo{db}
}

// Returns an array of all categories
func (r *categoryRepo) GetAll() (c []Category, err error) {
	var categories []Category

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&categories).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return categories, err
}

// Returns a category by id
func (r *categoryRepo) Get(id string) (c *Category, err error) {
	var category Category

	err = r.db.First(&category, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return &category, err
}
