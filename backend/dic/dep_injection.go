package dic

import (
	"country/domain/repo/category"
	"country/domain/repo/email"
	"country/domain/repo/region"
	"country/domain/repo/user"
	"country/infrastructure"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

// Dependency Injection Container
// Anything related to dependency injection goes here
// If I am using a DI framework might as well use it fully
var Container di.Container
var Builder *di.Builder

// Constants for accesing depedencies
const DB = "db"
const UserRepo = "user.repo"
const CategoryRepo = "category.repo"
const RegionRepo = "region.repo"
const EmailRepo = "email.repo"

// Initializes container and builder
func InitContainer() (di.Container, error) {
	builder, err := InitBuilder()
	if err != nil {
		return nil, err
	}

	Container = builder.Build()
	return Container, nil
}

// Initializes Builder and registers all the services
func InitBuilder() (*di.Builder, error) {
	Builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	err = RegisterServices(Builder)
	return Builder, err
}

func RegisterServices(builder *di.Builder) error {
	err := builder.Add(di.Def{
		Name: DB,
		Build: func(ctn di.Container) (interface{}, error) {
			return infrastructure.NewDB(), nil
		},
	})
	if err != nil {
		return err
	}

	err = builder.Add(di.Def{
		Name: UserRepo,
		Build: func(ctn di.Container) (interface{}, error) {
			return user.NewUserRepo(ctn.Get(DB).(*gorm.DB)), nil
		},
	})
	if err != nil {
		return err
	}
	err = builder.Add(di.Def{
		Name: CategoryRepo,
		Build: func(ctn di.Container) (interface{}, error) {
			return category.NewCategoryRepo(ctn.Get(DB).(*gorm.DB)), nil
		},
	})
	if err != nil {
		return err
	}
	err = builder.Add(di.Def{
		Name: RegionRepo,
		Build: func(ctn di.Container) (interface{}, error) {
			return region.NewRegionRepo(ctn.Get(DB).(*gorm.DB)), nil
		},
	})
	if err != nil {
		return err
	}

	err = builder.Add(di.Def{
		Name: EmailRepo,
		Build: func(ctn di.Container) (interface{}, error) {
			return email.NewEmailRepo(), nil
		},
	})
	if err != nil {
		return err
	}

	return nil
}
