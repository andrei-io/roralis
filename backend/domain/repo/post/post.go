// Contains everything related to interacting with posts
package post

import (
	"country/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type IPostRepo interface {
	GetAll() (c []entity.Post, err error)
	Get(id string) (c *entity.Post, err error)
	Update(id string, c *entity.Post) error
	Create(c *entity.Post) error
	Delete(id string) error
}

type PostRepo struct {
	db *gorm.DB
}

// Constructor function
func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) GetAll() (c []entity.Post, err error) {
	return nil, errors.New("Not implemented yet")
}

func (r *PostRepo) Get(id string) (c *entity.Post, err error) {
	var post entity.Post

	err = r.db.First(&post, id).Error

	return &post, err
}

func (r *PostRepo) Update(id string, c *entity.Post) error {
	return errors.New("Not implemented yet")
}

func (r *PostRepo) Create(c *entity.Post) error {
	return errors.New("Not implemented yet")
}

func (r *PostRepo) Delete(id string) error {
	return errors.New("Not implemented yet")
}
