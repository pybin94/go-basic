package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("[Error] Not Found")
	errCantUpdate = errors.New("[Error] Can't update non-existing word")
	errCantDELETE = errors.New("[Error] Can't delete non-existing word")
	errWordExists = errors.New("[Error] That word is exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}

	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
