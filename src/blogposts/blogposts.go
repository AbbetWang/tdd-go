package blogposts

import (
	"io/fs"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(fs fs.FS) []Post {
	return nil
}
