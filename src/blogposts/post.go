package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title, Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	scanner.Scan()
	titleLine := scanner.Text()
	scanner.Scan()
	descriptionLine := scanner.Text()
	return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil
}
