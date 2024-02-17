package maps

import (
	"errors"
)

/*
	NEVER declare map like this 
	var dict map[string]string

	INSTEAD use this
	var dict = make(map[string]string)
	var dict = map[string]string{}
*/

var ErrorNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("key already exists")

func (d *Dictionary) Search(searchWord string) (string, error) {
	got, ok := d.Dict[searchWord]
	if !ok {
		return "", ErrorNotFound
	}
	return got, nil
}

func (d *Dictionary) Add(key, value string) (error){
	_, err := d.Search(key)
	if err == nil {
		return ErrWordExists
	}
	d.Dict[key] = value
	return nil
}

type Dictionary struct {
	Dict map[string]string
}

