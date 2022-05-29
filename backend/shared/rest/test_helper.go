package rest

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

type KV struct {
	Key   string
	Value string
}

type TestHttpConfig struct {
	Header      http.Header
	QueryParams []KV
	Body        []byte
}

func NewMockGinContext(config *TestHttpConfig) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	if config == nil {
		return c, w
	}

	req := &http.Request{
		URL:    &url.URL{},
		Header: config.Header,
	}
	if config.Body != nil {
		bodyReadCloser := io.NopCloser(bytes.NewReader(config.Body))
		req.Body = bodyReadCloser
	}

	if config.QueryParams != nil {
		q := req.URL.Query()
		for _, query := range config.QueryParams {
			q.Add(query.Key, query.Value)
		}
		// must set this, since under the hood c.BindQuery calls
		// `req.URL.Query()`, which calls `ParseQuery(u.RawQuery)`
		req.URL.RawQuery = q.Encode()
	}

	// finally set the request to the gin context
	c.Request = req
	return c, w
}
