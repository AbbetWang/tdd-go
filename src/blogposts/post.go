package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title, Description string
	Tags               []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}, nil
}

func trimSpace(tags []string) []string {
	for i, tag := range tags {
		tags[i] = strings.TrimSpace(tag)
	}
	return tags
}
