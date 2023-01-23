package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

// var errSearch = errors.New("NOT in the list")
// var errAdd = errors.New("Word in the list already")
// var errUpdate = errors.New("Can't Update")

var (
	errSearch = errors.New("NOT in the list")
	errAdd    = errors.New("Word in the list already")
	errUpdate = errors.New("Can't Update")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //map returns value, ok (which is bool)
	if exists {
		return value, nil
	}
	return "", errSearch
}

// Add a word and its definition
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errSearch:
		d[word] = def
	case nil:
		return errAdd
	}
	return nil
}

// Update a word definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errSearch:
		return errUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}