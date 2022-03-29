package post

import "backend/roralis/domain/repo/post"

type PostController struct {
	repo     post.PostRepo
	tokenKey string
}

func NewPostController(r post.PostRepo, t string) PostController {
	return PostController{repo: r, tokenKey: t}
}
