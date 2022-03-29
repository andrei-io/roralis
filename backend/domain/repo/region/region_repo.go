// Contains everything related to working with regions: repo
package region

import (
	"backend/roralis/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type RegionRepo interface {
	GetAll() (re []entity.Region, err error)
	Get(id string) (re *entity.Region, err error)
	Update(id string, re *entity.Region) error
	Create(re *entity.Region) error
	Delete(id string) error
}

type regionRepo struct {
	db *gorm.DB
}

// Check interface at compile time
var _ RegionRepo = (*regionRepo)(nil)

// Constructor function
func NewRegionRepo(db *gorm.DB) *regionRepo {
	return &regionRepo{db}
}

// Returns an array of all regions
func (r *regionRepo) GetAll() (re []entity.Region, err error) {
	var regions []entity.Region

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&regions).Error

	return regions, err
}

// Get one region by id
func (r *regionRepo) Get(id string) (re *entity.Region, err error) {
	var region entity.Region

	err = r.db.First(&region, id).Error

	return &region, err
}

func (r *regionRepo) Update(id string, re *entity.Region) error {
	return errors.New("Modifying regions is not allowed")
}

func (r *regionRepo) Create(re *entity.Region) error {
	return errors.New("Modifying regions is not allowed")
}

func (r *regionRepo) Delete(id string) error {
	return errors.New("Modifying regions is not allowed")
}
