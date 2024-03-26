package main

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("could not add the word it already exists")
)

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}
func (e DictionaryErr) Error() string {
	return string(e)
}
