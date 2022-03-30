package post

import (
	"errors"

	"gorm.io/gorm"
)

type PostRepo interface {
	GetAll(offset int, limit int, newest bool) (c []Post, err error)
	Get(id string) (p *Post, err error)
	Update(id string, p *Post) error
	Create(p *Post) error
	Delete(id string) error
}

type postRepo struct {
	db *gorm.DB
}

// Check interface at compile time
var _ PostRepo = (*postRepo)(nil)

// Constructor function
func NewPostRepo(db *gorm.DB) *postRepo {
	return &postRepo{db: db}
}

// Gets one post by id
func (r *postRepo) Get(id string) (p *Post, err error) {
	var post Post

	err = r.db.First(&post, id).Error

	return &post, err
}

// Gets all the posts with pagination and if set orderds by date created
func (r *postRepo) GetAll(offset int, limit int, newest bool) (c []Post, err error) {
	var posts []Post
	var order string

	if newest {
		order = "created_at desc"
	} else {
		order = "id"
	}

	err = r.db.Offset(offset).Limit(limit).Order(order).Find(&posts).Error

	return posts, err
}

func (r *postRepo) Update(id string, p *Post) error {
	return errors.New("Not implemented yet")
}

func (r *postRepo) Create(p *Post) error {
	err := r.db.Create(&p).Error
	return err
}

func (r *postRepo) Delete(id string) error {
	return errors.New("Not implemented yet")
}
