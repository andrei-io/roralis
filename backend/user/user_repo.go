package user

import (
	"strconv"
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
	var user []User

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&user).Error

	return user, err
}

func (r *userRepo) Get(id string) (u *User, err error) {
	var user User

	err = r.db.First(&user, id).Error

	return &user, err
}

func (r *userRepo) Update(id string, u *User) error {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	u.ID = uid
	// More complicated than just db.Save(&u), but this way we can return a 404
	operation := r.db.Model(&User{}).Where("id = ?", id).Updates(&u)
	if operation.Error != nil {
		message := operation.Error.(*pgconn.PgError).Message
		if strings.Contains(message, "duplicate key value violates unique constraint") {
			return operation.Error
		}
		if operation.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}

	return nil
}

func (r *userRepo) Create(u *User) error {
	err := r.db.Create(&u).Error
	return err
}

func (r *userRepo) Delete(id string) error {
	operation := r.db.Delete(&User{}, id)
	if operation.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return operation.Error
}

func (r *userRepo) GetByEmail(email string) (u *User, err error) {
	var user User

	err = r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
