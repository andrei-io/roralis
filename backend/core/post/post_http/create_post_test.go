package posthttp_test

import (
	"backend/roralis/core/jwt"
	"backend/roralis/core/post"
	posthttp "backend/roralis/core/post/post_http"
	"backend/roralis/core/user"
	"backend/roralis/shared/rest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreatePost(t *testing.T) {
	testUser := user.User{
		ID:       1,
		Name:     "FirstUser",
		Email:    "first@example.com",
		Password: "DefinetlyHashedPassword",
		Verified: true,
		Role:     5,
	}

	mockRepo := postRepoMock{
		notFoundError: false,
		data: []post.Post{
			{ID: 1, UserID: 1, Title: "First Post"},
			{ID: 2, UserID: 2, Title: "Second Post"},
		},
	}

	postController := posthttp.NewPostController(&mockRepo, "testing")

	// Error on invalid permissions
	c, w := rest.NewMockGinContext(nil)
	c.Set("testing", &jwt.JWTClaims{
		Name:     testUser.Name,
		ID:       testUser.ID,
		Verified: testUser.Verified,
		// 5 is the minimiun permission level
		Role: 1,
	})

	postController.Create(c)
	if w.Code != http.StatusForbidden {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusForbidden, w.Code, w.Body.String())
	}

	// Succesfully created

	body, err := json.Marshal(&post.Post{
		Title:      "Testing",
		RegionID:   1,
		CategoryID: 1,
	})
	if err != nil {
		t.Errorf("Error on marshalling json: %+v", err)
	}
	c, w = rest.NewMockGinContext(&rest.TestHttpConfig{Body: body})
	c.Set("testing", &jwt.JWTClaims{
		Name:     testUser.Name,
		ID:       testUser.ID,
		Verified: testUser.Verified,
		// 5 is the minimiun permission level
		Role: testUser.Role,
	})
	postController.Create(c)
	if w.Code != http.StatusOK {
		t.Errorf("Got wrong http code, wanted %v, got %v, \nJSON: %+v", http.StatusOK, w.Code, w.Body.String())
	}

}
