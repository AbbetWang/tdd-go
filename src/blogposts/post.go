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

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()

	}
	title := readLine()[7:]
	description := readLine()[13:]
	return Post{Title: title, Description: description}, nil
}
