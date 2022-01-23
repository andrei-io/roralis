package category

import (
	"country/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type ICategoryRepo interface {
	GetAll() (users []entity.Category, err error)
	Get(id string) (u *entity.Category, err error)
	Update(id string, u *entity.Category) error
	Create(u *entity.Category) error
	Delete(id string) error
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

// Returns an array of all users
func (r *CategoryRepo) GetAll() (users []entity.Category, err error) {
	var categories []entity.Category

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&categories).Error

	return categories, err
}

func (r *CategoryRepo) Get(id string) (u *entity.Category, err error) {
	var user entity.Category

	err = r.db.First(&user, id).Error

	return &user, err
}

func (r *CategoryRepo) Update(id_raw string, u *entity.Category) error {
	return errors.New("Modifying categories is not allowed")
}

func (r *CategoryRepo) Create(u *entity.Category) error {
	return errors.New("Modifying categories is not allowed")
}

func (r *CategoryRepo) Delete(id_raw string) error {
	return errors.New("Modifying categories is not allowed")
}
