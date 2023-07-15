package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Word not found")
var errKeyExists = errors.New("Word already exists")

func (dictionary_ Dictionary) Search(word_ string) (string, error) {
	definition, exists := dictionary_[word_]
	if exists {
		return definition, nil
	} else {
		return "", errNotFound
	}
}

// Add a word to the dictionary
func (dictionary_ Dictionary) Add(word_, def_ string) error {
	_, err := dictionary_.Search(word_)

	switch err {
	case errNotFound:
		dictionary_[word_] = def_
	case nil:
		return errKeyExists
	}

	return nil
}
