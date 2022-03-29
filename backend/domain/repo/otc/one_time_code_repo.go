// Contains everything related to working with OneTimeCodes
package otc

import (
	"backend/roralis/domain/entity"
	"errors"
	"time"

	"gorm.io/gorm"
)

type OTCRepo interface {
	Set(key uint64, value string, expire_minutes int64) error
	Get(key uint64) (string, error)
	Delete(key string) error
}

type otcRepo struct {
	db *gorm.DB
}

// Check interface at compile time
var _ OTCRepo = (*otcRepo)(nil)

// Constructor function
func NewOTCRepo(db *gorm.DB) *otcRepo {
	return &otcRepo{db}
}

// Sets a (key, value) pair
func (k *otcRepo) Set(key uint64, value string, expire_minutes int64) error {
	otc := entity.OneTimeCode{
		UserID: key,
		Code:   value,
		Active: true,
		Expire: time.Now().Add(time.Minute * time.Duration(expire_minutes)),
	}
	err := k.db.Create(&otc).Error
	return err

}

// Gets a (key, value) pair
func (k *otcRepo) Get(key uint64) (string, error) {
	var otc entity.OneTimeCode

	err := k.db.Where("user_id = ?", key).First(&otc).Error
	if err != nil {
		return "", err
	}
	if time.Now().After(otc.Expire) {
		return "", errors.New("Expiration code has expired")
	}
	return otc.Code, nil
}

func (k *otcRepo) Delete(key string) error {
	return errors.New("Not implemented yet")
}
