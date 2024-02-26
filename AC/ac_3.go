package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func changeemail(h int, new_email string, contatos [5]Contato) [5]Contato {
	if contatos[h].Email != "" {
		contatos[h].Email = new_email
	}
	return contatos
}

func adicionar_Contato(nome string, email string, contatos [5]Contato) [5]Contato {
	novo_contato := Contato{Nome: nome, Email: email}
	for i := range contatos {
		if contatos[i].Email == "" || contatos[i].Nome == "" {
			contatos[i] = novo_contato
			break
		}
	}
	return contatos
}

func excluir_Contato(contatos [5]Contato) [5]Contato {
	for i := len(contatos) - 1; i >= 0; i-- {
		if contatos[i].Email != "" && contatos[i].Nome != "" {
			contatos[i].Email = ""
			contatos[i].Nome = ""
			break
		}
	}
	return contatos
}

func main() {
	var contatos [5]Contato

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Adicionar contato")
		fmt.Println("2. Excluir o último contato")
		fmt.Println("3. Sair")

		scanner.Scan()
		opcao := scanner.Text()

		switch opcao {
		case "1":
			fmt.Println("Digite o nome do contato:")
			scanner.Scan()
			nome := scanner.Text()

			fmt.Println("Digite o email do contato:")
			scanner.Scan()
			email := scanner.Text()

			contatos = adicionar_Contato(nome, email, contatos)
			fmt.Println("Contato adicionado com sucesso!")
		case "2":
			contatos = excluir_Contato(contatos)
			fmt.Println("Último contato excluído com sucesso!")
		case "3":
			fmt.Println("Saindo...")
			fmt.Println("Contatos: ")
			for _, c:= range contatos{
				fmt.Println(c.Nome, c.Email)
			}
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

	}
}
