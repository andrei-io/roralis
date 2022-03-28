// Contains everything related to working with the Category entity: repo
package category

import (
	"backend/roralis/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type ICategoryRepo interface {
	GetAll() (c []entity.Category, err error)
	Get(id string) (c *entity.Category, err error)
	Update(id string, c *entity.Category) error
	Create(c *entity.Category) error
	Delete(id string) error
}

type CategoryRepo struct {
	db *gorm.DB
}

// Constructr funcion
func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

// Returns an array of all categories
func (r *CategoryRepo) GetAll() (c []entity.Category, err error) {
	var categories []entity.Category

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&categories).Error

	return categories, err
}

// Returns a category by id
func (r *CategoryRepo) Get(id string) (c *entity.Category, err error) {
	var category entity.Category

	err = r.db.First(&category, id).Error

	return &category, err
}

func (r *CategoryRepo) Update(id string, c *entity.Category) error {
	return errors.New("Modifying categories is not allowed")
}

func (r *CategoryRepo) Create(c *entity.Category) error {
	return errors.New("Modifying categories is not allowed")
}

func (r *CategoryRepo) Delete(id string) error {
	return errors.New("Modifying categories is not allowed")
}
