package maps

type Dictionary map[string]string

// We could reuse ErrNotFound and not add a new error. However, it is often better to have a precise error for cruds.
const (
	ErrNotFound                = DictionaryErr("could not find the word you were looking for")
	ErrWordExists              = DictionaryErr("cannot add word because it already exists")
	ErrUpdatingNonExistingWord = DictionaryErr("cannot perform update on word because it does not exist")
	ErrDeletingNonExistingWord = DictionaryErr("cannot perform delete on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrUpdatingNonExistingWord
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrDeletingNonExistingWord
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
