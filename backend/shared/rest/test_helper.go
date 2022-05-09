package rest

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewMockGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}
