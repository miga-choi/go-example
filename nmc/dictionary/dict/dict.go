package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Word not found")

func (dictionary_ Dictionary) Search(word_ string) (string, error) {
	definition, exists := dictionary_[word_]
	if exists {
		return definition, nil
	} else {
		return "", errNotFound
	}
}
