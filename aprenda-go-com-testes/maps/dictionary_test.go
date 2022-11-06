package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"teste": "isso é apenas um teste"}

	t.Run("Knowledge word", func(t *testing.T) {
		result, _ := dictionary.Search("teste")
		expect := "isso é apenas um teste"

		compareStrings(t, result, expect)
	})

	t.Run("No Knowledge word", func(t *testing.T) {
		_, err := dictionary.Search("desconhecida")

		compareError(t, err, ErrNotFound)
	})

}

func compareStrings(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("result '%s', expect '%s', dado '%s'", result, expect, "teste")
	}
}

func compareError(t *testing.T, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("result erro '%s', expect '%s'", result, expect)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "teste"
		definition := "isso é apenas um teste"

		err := dictionary.Add(word, definition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, definition)
	})

	t.Run("existent word", func(t *testing.T) {
		word := "teste"
		definition := "isso é apenas um teste"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "teste novo")

		compareError(t, err, ErrWordAlreadyExistent)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existent word", func(t *testing.T) {
		word := "teste"
		definition := "isso é apenas um teste"

		dictionary := Dictionary{word: definition}

		newDefinition := "nova definição"

		err := dictionary.Update(word, newDefinition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "teste"
		definition := "isso é apenas um teste"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		compareError(t, err, ErrWorldNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "teste"
	dictionary := Dictionary{word: "definição de teste"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("espera-se que '%s' seja deletado", word)
	}
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should has find the added word", err)
	}

	if definition != result {
		t.Errorf("result '%s', expect '%s'", result, definition)
	}
}
