package user

import (
	"country/domain/entity"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type IUserRepo interface {
	GetAll() (users []entity.User, err error)
	Get(id string) (u *entity.User, err error)
	GetByEmail(email string) (u *entity.User, err error)
	Update(id string, u *entity.User) error
	Create(u *entity.User) error
	Delete(id string) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// Returns an array of all users
func (r *UserRepo) GetAll() (users []entity.User, err error) {
	var user []entity.User

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&user).Error

	return user, err
}

func (r *UserRepo) Get(id string) (u *entity.User, err error) {
	var user entity.User

	err = r.db.First(&user, id).Error

	return &user, err
}

func (r *UserRepo) Update(id string, u *entity.User) error {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	u.ID = uid
	// More complicated than just db.Save(&u), but this way we can return a 404
	operation := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(&u)
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

func (r *UserRepo) Create(u *entity.User) error {
	err := r.db.Create(&u).Error
	return err
}

func (r *UserRepo) Delete(id string) error {
	operation := r.db.Delete(&entity.User{}, id)
	if operation.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return operation.Error
}

func (r *UserRepo) GetByEmail(email string) (u *entity.User, err error) {
	var user entity.User

	err = r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
