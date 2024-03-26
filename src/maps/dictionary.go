package main

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("could not add the word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("could not update the word which does not exist")
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
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
