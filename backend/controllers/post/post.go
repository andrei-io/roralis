package post

import "backend/roralis/domain/repo/post"

type PostController struct {
	repo post.PostRepo
}

func NewPostController(r post.PostRepo) PostController {
	return PostController{repo: r}
}
