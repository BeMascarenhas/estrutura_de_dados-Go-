package main

import "fmt"

type Nó struct {
	valor int
	prox  *Nó
}

type Fila struct {
	tamanho int
	início  *Nó
	fim     *Nó
}

func criarFila() *Fila {
	return &Fila{}
}

func (f *Fila) enfileirar(dado int) {
	temp := &Nó{valor: dado}

	if f.início == nil {
		f.início = temp
		f.fim = temp
		f.tamanho = 1
		return
	}

	f.fim.prox = temp
	f.fim = temp
	f.tamanho++
}

func (f *Fila) desenfileirar() {
	if f.início == nil {
		fmt.Println("A fila está vazia.")
		return
	}

	f.início = f.início.prox
	f.tamanho--

	if f.início == nil {
		f.fim = nil
	}
}

func (f *Fila) percorrerFila() {
	atual := f.início
	fmt.Print("Elementos da fila: ")
	for atual != nil {
		fmt.Printf("%d ", atual.valor)
		atual = atual.prox
	}
	fmt.Println()
}

func main() {
	fila := criarFila()
	fila.enfileirar(10)
	fila.enfileirar(20)
	fila.enfileirar(30)
	fila.percorrerFila()
	fila.desenfileirar()
	fila.percorrerFila()
}
