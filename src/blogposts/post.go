package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title, Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	return Post{Title: title, Description: description}, nil
}
