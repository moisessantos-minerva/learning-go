
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func verificarSenha(senha string) (bool, bool, bool) {
	tamanhoValido := utf8.RuneCountInString(senha) >= 8
	temNumero := false
	temMaiuscula := false

	for _, caractere := range senha {
		if unicode.IsDigit(caractere) {
			temNumero = true
		}

		if unicode.IsUpper(caractere) {
			temMaiuscula = true
		}
	}

	return tamanhoValido, temNumero, temMaiuscula
}

func senhaConfirmada(senha, confirmacao string) bool {
	return senha == confirmacao
}

func mostrarResultado(criterio string, valido bool) {
	if valido {
		fmt.Println("[OK]", criterio)
	} else {
		fmt.Println("[X]", criterio)
	}
}

func lerEntrada(leitor *bufio.Reader) (string, error) {
	entrada, err := leitor.ReadString('\n')
	if err != nil && len(entrada) == 0 {
		return "", err
	}
	return strings.TrimRight(entrada, "\r\n"), nil
}

func main() {
	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("=== Verificador de Senhas em Go ===")

	fmt.Print("Digite uma senha: ")
	senha, err := lerEntrada(leitor)
	if err != nil {
		fmt.Println("Erro ao ler senha:", err)
		return
	}

	fmt.Print("Confirme a senha: ")
	confirmacao, err := lerEntrada(leitor)
	if err != nil {
		fmt.Println("Erro ao ler confirmacao:", err)
		return
	}

	tamanho, numero, maiuscula := verificarSenha(senha)
	confere := senhaConfirmada(senha, confirmacao)

	fmt.Println("\nResultado da verificacao:")
	mostrarResultado("Pelo menos 8 caracteres", tamanho)
	mostrarResultado("Contem numero", numero)
	mostrarResultado("Contem letra maiuscula", maiuscula)
	mostrarResultado("Senhas iguais", confere)

	if tamanho && numero && maiuscula && confere {
		fmt.Println("\nSenha atende aos criterios!")
	} else {
		fmt.Println("\nSenha nao atende aos criterios.")
	}
}
