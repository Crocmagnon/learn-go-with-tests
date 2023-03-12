package _7_maps

// Dictionary stores words and their definitions.
type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search gets a word from the dictionary and returns its definition.
// It returns ErrNotFound if the word wasn't found in the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a word and its definition to the dictionary,
// returning ErrWordExists if the word already exists in the dictionary.
// In this case, use Update.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update updates the definition of a word in the dictionary.
// It returns ErrWordDoesNotExist when the word doesn't already exist.
// In this case, use Add.
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

// Delete deletes a word and its definition from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
