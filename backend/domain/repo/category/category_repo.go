// Contains everything related to working with the Category entity: repo
package category

import (
	"backend/roralis/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	GetAll() (c []entity.Category, err error)
	Get(id string) (c *entity.Category, err error)
	Update(id string, c *entity.Category) error
	Create(c *entity.Category) error
	Delete(id string) error
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
func (r *categoryRepo) GetAll() (c []entity.Category, err error) {
	var categories []entity.Category

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&categories).Error

	return categories, err
}

// Returns a category by id
func (r *categoryRepo) Get(id string) (c *entity.Category, err error) {
	var category entity.Category

	err = r.db.First(&category, id).Error

	return &category, err
}

func (r *categoryRepo) Update(id string, c *entity.Category) error {
	return errors.New("Modifying categories is not allowed")
}

func (r *categoryRepo) Create(c *entity.Category) error {
	return errors.New("Modifying categories is not allowed")
}

func (r *categoryRepo) Delete(id string) error {
	return errors.New("Modifying categories is not allowed")
}
