
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

func mostrarResultado(criterio string, valido bool) {
	if valido {
		fmt.Println("[OK]", criterio)
	} else {
		fmt.Println("[X]", criterio)
	}
}

func main() {
	fmt.Println("=== Verificador de Senhas em Go ===")
	fmt.Print("Digite uma senha: ")

	leitor := bufio.NewReader(os.Stdin)
	senha, err := leitor.ReadString('\n')

	if err != nil && len(senha) == 0 {
		fmt.Println("Erro ao ler a senha:", err)
		return
	}

	senha = strings.TrimRight(senha, "\r\n")

	tamanho, numero, maiuscula := verificarSenha(senha)

	fmt.Println("\nResultado da verificacao:")
	mostrarResultado("Pelo menos 8 caracteres", tamanho)
	mostrarResultado("Contem numero", numero)
	mostrarResultado("Contem letra maiuscula", maiuscula)

	if tamanho && numero && maiuscula {
		fmt.Println("\nSenha atende aos criterios!")
	} else {
		fmt.Println("\nSenha nao atende aos criterios.")
	}
}
