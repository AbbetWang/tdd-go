package main

type Dictionary map[string]string

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

func (d Dictionary) Search(word string) string {
	return d[word]
}
