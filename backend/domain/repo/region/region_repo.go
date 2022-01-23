package region

import (
	"country/domain/entity"
	"errors"

	"gorm.io/gorm"
)

type IRegionRepo interface {
	GetAll() (re []entity.Region, err error)
	Get(id string) (re *entity.Region, err error)
	Update(id string, re *entity.Region) error
	Create(re *entity.Region) error
	Delete(id string) error
}

type RegionRepo struct {
	db *gorm.DB
}

func NewRegionRepo(db *gorm.DB) *RegionRepo {
	return &RegionRepo{db}
}

// Returns an array of all regions
func (r *RegionRepo) GetAll() (re []entity.Region, err error) {
	var regions []entity.Region

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&regions).Error

	return regions, err
}

// Get one region by id
func (r *RegionRepo) Get(id string) (re *entity.Region, err error) {
	var region entity.Region

	err = r.db.First(&region, id).Error

	return &region, err
}

func (r *RegionRepo) Update(id string, re *entity.Region) error {
	return errors.New("Modifying regions is not allowed")
}

func (r *RegionRepo) Create(re *entity.Region) error {
	return errors.New("Modifying regions is not allowed")
}

func (r *RegionRepo) Delete(id string) error {
	return errors.New("Modifying regions is not allowed")
}
