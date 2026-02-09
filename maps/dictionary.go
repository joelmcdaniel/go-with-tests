package maps

type Dictionary map[string]string

const (
	ErrNotFound     = DictionaryErr("could not find word in the dictionary")
	ErrWordExists   = DictionaryErr("word already exists in the dictionary, cannot add")
	ErrDoesNotExist = DictionaryErr("word doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if definition, ok := d[word]; ok {
		return definition, nil
	}

	return "", ErrNotFound
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
		return ErrDoesNotExist
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
		return ErrDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
