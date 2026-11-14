
package main

import "testing"

func TestSenhaValidaComConfirmacao(t *testing.T) {
	tamanho, numero, maiuscula := verificarSenha("Teste123")
	confere := senhaConfirmada("Teste123", "Teste123")

	if !(tamanho && numero && maiuscula && confere) {
		t.Error("Senha valida com confirmacao igual deveria ser aceita")
	}
}

func TestConfirmacaoDiferente(t *testing.T) {
	if senhaConfirmada("Teste123", "Teste124") {
		t.Error("Confirmacao diferente deveria ser rejeitada")
	}
}

func TestSenhaCurta(t *testing.T) {
	tamanho, _, _ := verificarSenha("Ab1")

	if tamanho {
		t.Error("Senha com menos de 8 caracteres deveria ser rejeitada")
	}
}

func TestSenhaSemMaiuscula(t *testing.T) {
	_, _, maiuscula := verificarSenha("teste123")

	if maiuscula {
		t.Error("Senha sem letra maiuscula deveria ser rejeitada")
	}
}

func TestConfirmacaoDiferenciaMaiusculas(t *testing.T) {
	if senhaConfirmada("Teste123", "teste123") {
		t.Error("Confirmacao deve diferenciar maiusculas e minusculas")
	}
}
