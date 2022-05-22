package region

import (
	"backend/roralis/shared/repo"
	"errors"

	"gorm.io/gorm"
)

type RegionRepo interface {
	GetAll() (re []Region, err error)
	Get(id string) (re *Region, err error)
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
func (r *regionRepo) GetAll() (re []Region, err error) {
	var regions []Region

	// Will panic on fail,but gin has a recovery middleware
	err = r.db.Find(&regions).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return regions, err
}

// Get one region by id
func (r *regionRepo) Get(id string) (re *Region, err error) {
	var region Region

	err = r.db.First(&region, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, repo.ErrRecordNotFound
	}

	return &region, err
}
