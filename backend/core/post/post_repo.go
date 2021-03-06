package post

import (
	"backend/roralis/shared/repo"
	"errors"

	"gorm.io/gorm"
)

type PostRepo interface {
	GetAll(offset int, limit int, newest bool) ([]Post, error)
	Get(id string) (*Post, error)
	GetByUserID(id string) ([]Post, error)
	Create(p *Post) error
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
func (r *postRepo) Get(id string) (*Post, error) {
	var post Post

	err := r.db.First(&post, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return &post, err
}

// Gets all the posts with pagination and if set orderds by date created
func (r *postRepo) GetAll(offset int, limit int, newest bool) ([]Post, error) {
	var posts []Post
	var order string

	if newest {
		order = "created_at desc"
	} else {
		order = "id"
	}

	err := r.db.Offset(offset).Limit(limit).Order(order).Find(&posts).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return posts, err
}

func (r *postRepo) GetByUserID(id string) ([]Post, error) {
	var posts []Post

	err := r.db.Where("user_id = ?", id).Order("created_at desc").Find(&posts).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return posts, err
}

func (r *postRepo) Create(p *Post) error {
	err := r.db.Create(&p).Error
	return err
}
