# Aprendendo Go — Verificador de Senhas

Projeto desenvolvido como parte da atividade **Aprender com IA sem terceirizar o aprendizado**, da Soul/Minerva, com o objetivo de conhecer os fundamentos da linguagem Go.

## Funcionalidades

- Receber uma senha pelo terminal.
- Verificar se possui pelo menos 8 caracteres.
- Verificar se contém pelo menos um número.
- Verificar se contém pelo menos uma letra maiúscula.
- Exibir os critérios atendidos e não atendidos.

## Como executar

É necessário ter o Go instalado.

```bash
go run .
```

## Conceitos explorados

- Estrutura de um programa em Go (`package main`).
- Declaração de variáveis com `:=`.
- Funções e múltiplos valores de retorno.
- Condicionais e laços `for`.
- Manipulação de strings e caracteres Unicode.

## Testes manuais

| Entrada | Resultado esperado |
|---|---|
| `Teste123` | Atende aos critérios |
| `teste123` | Falta letra maiúscula |
| `Testeabc` | Falta número |
| `Ab1` | Menos de 8 caracteres |
| `12345678` | Falta letra maiúscula |

Todos os cinco cenários foram executados e apresentaram os resultados esperados.

## Referências oficiais

- [Go — For statements with range clause](https://go.dev/ref/spec#For_range)
- [Go — Package unicode](https://pkg.go.dev/unicode)

## Observação

Este é um projeto educacional, não uma implementação completa de políticas de segurança de senhas.