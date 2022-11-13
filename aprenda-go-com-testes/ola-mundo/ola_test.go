package main

import "testing"

// func TestHello(t *testing.T) {
// 	result := Hello()

// 	expected := "Olá mundo!"

// 	if result != expected {
// 		t.Errorf("resultado '%s', esperado '%s'", result, expected)
// 	}
// }

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		// Se não for chamado, se o teste falhar, irá apontar para linha que falhou e não para a chamada do helper
		t.Helper()
		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	}
	t.Run("Diz olá para as pessoas", func(t *testing.T) {
		result := Hello("Chris", "")
		expected := "Olá, Chris"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("Diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		result := Hello("", "")
		expected := "Olá, Mundo"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("Em espanhol", func(t *testing.T) {
		result := Hello("Fulano", "espanhol")
		expected := "Hola, Fulano"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("Em Francês", func(t *testing.T) {
		result := Hello("Fulano", "frances")
		expected := "Bonjour, Fulano"

		verifyCorrectMessage(t, result, expected)
	})
}
