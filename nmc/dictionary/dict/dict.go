package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Word not found")
	errKeyExists  = errors.New("Word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

// Search a word from dictionary
func (dictionary_ Dictionary) Search(word_ string) (string, error) {
	definition, exists := dictionary_[word_]
	if exists {
		return definition, nil
	} else {
		return "", errNotFound
	}
}

// Add a word to the dictionary
func (dictionary_ Dictionary) Add(word_, definition_ string) error {
	_, err := dictionary_.Search(word_)

	switch err {
	case errNotFound:
		dictionary_[word_] = definition_
	case nil:
		return errKeyExists
	}

	return nil
}

// Update a word
func (dictionary_ Dictionary) Update(word_, definition_ string) error {
	_, err := dictionary_.Search(word_)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		dictionary_[word_] = definition_
	}

	return nil
}

// Delete a word
func (dictionary_ Dictionary) Delete(word_ string) {
	delete(dictionary_, word_)
}
