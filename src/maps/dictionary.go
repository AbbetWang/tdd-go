package main

import "errors"

type Dictionary map[string]string

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
