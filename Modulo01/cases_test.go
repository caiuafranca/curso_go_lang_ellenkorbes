package main

import (
	"fmt"
	"testing"
)

// Aprendendo GO com testes
// https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/ola-mundo

func TestVariaveis(t *testing.T) {
	x := 42.1
	y := "James Bold"
	z := true
	testeInt := fmt.Sprintf("%T", x)
	testeString := fmt.Sprintf("%T", y)
	testeBool := fmt.Sprintf("%T", z)

	if testeInt != "int" {
		t.Error("Ocorreu um erro!!!, verifique o tipo do dado, saida: ", x)
	}
	if testeString != "string" {
		t.Error("Ocorreu um erro!!!, verifique o tipo do dado, saida", y)
	}
	if testeBool != "bool" {
		t.Error("Ocorreu um erro!!!, verifique o tipo do dado, saida", z)
	}
}
