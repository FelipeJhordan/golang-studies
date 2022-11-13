package main

const (
	ErrNotFound            = ErrDictionary("not is possible find the word that you serch")
	ErrWordAlreadyExistent = ErrDictionary("not is possible add a world because his already registered.")
	ErrWorldNotFound       = ErrDictionary("not is possible update the world because his not exists")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}


type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]

	if !exist {
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
		return ErrWordAlreadyExistent
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWorldNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
