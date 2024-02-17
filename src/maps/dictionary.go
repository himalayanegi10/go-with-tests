package maps

import (
	"errors"
)

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (d *Dictionary) Search(searchWord string) (string, error) {
	got, ok := d.Dict[searchWord]
	if !ok {
		return "", ErrorNotFound
	}
	return got, nil
}

type Dictionary struct {
	Dict map[string]string
}

