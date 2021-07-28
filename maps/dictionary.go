package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	var m map[string]string
	fmt.Printf("%T %p %v %s====", m, m, m, m["word"])
	d[word] = definition
}
