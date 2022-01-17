package user

import (
	"country/dic"
	"country/domain/entity"
	"country/domain/repo/user"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

// Gin controller for GET /users
func ReadAll(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)

	users, err := repo.GetAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

// Gin controller for GET /users/:id
func ReadOne(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	id := c.Param("id")

	u, err := repo.Get(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, entity.NotFoundError)
		return
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

// Gin controller for PUT /users/:id
func Update(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	var json entity.User
	id := c.Param("id")

	// Validate request form
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Message: err.Error()})
		return
	}

	// Update in DB. Will error out when invalid
	err := repo.Update(id, &json)
	if err != nil {
		err := err.(*pgconn.PgError)

		message := err.Message
		if strings.Contains(message, "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, entity.NewDuplicateEntityErrorResponse(err.ConstraintName))
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, entity.NotFoundError)
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, entity.SuccesResponse)
}

// Gin controller for POST /users/
func Create(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	var json entity.User

	// Validate request form
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, entity.Response{Message: err.Error()})
		return
	}

	// Create in db. Will error out when invalid
	err := repo.Create(&json)
	if err != nil {
		err := err.(*pgconn.PgError)
		message := err.Message
		if strings.Contains(message, "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, entity.NewDuplicateEntityErrorResponse(err.ConstraintName))
			return
		} else {
			c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, entity.SuccesResponse)

}

// Gin controller for DELETE /users/:id
func Delete(c *gin.Context) {
	repo := dic.Container.Get(dic.UserRepo).(user.IUserRepo)
	id := c.Param("id")

	err := repo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, entity.Response{Message: err.Error()})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, entity.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity.SuccesResponse)

}
