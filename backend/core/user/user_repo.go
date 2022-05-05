package user

import (
	"backend/roralis/shared/repo"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetAll() (users []User, err error)
	Get(id string) (u *User, err error)
	GetByEmail(email string) (u *User, err error)
	Update(id string, u *User) error
	Create(u *User) error
	Delete(id string) error
}

type userRepo struct {
	db *gorm.DB
}

// Check interface at compile time
var _ UserRepo = (*userRepo)(nil)

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

// Returns an array of all users
func (r *userRepo) GetAll() (users []User, err error) {
	return nil, repo.ErrNotImplementedYet
}

func (r *userRepo) Get(id string) (u *User, err error) {
	var user User

	err = r.db.First(&user, id).Error

	return &user, err
}

func (r *userRepo) Update(id string, u *User) error {
	return repo.ErrNotImplementedYet
}

func (r *userRepo) Create(u *User) error {
	err := r.db.Create(&u).Error
	if err != nil {
		err := err.(*pgconn.PgError)
		fmt.Println(err.Message)
		if strings.Contains(err.Message, "duplicate key value violates unique constraint") {
			return repo.ErrEmailTaken
		}
	}
	return nil
}

func (r *userRepo) Delete(id string) error {
	return repo.ErrNotImplementedYet
}

func (r *userRepo) GetByEmail(email string) (u *User, err error) {
	var user User

	err = r.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return &user, err
}
