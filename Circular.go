package main

import "fmt"


type No struct {
	Value int
	Next  *No
}

type listaCircular struct {
	Cabeca *No
}

func (lista *listaCircular) Exibir() {
	if lista.Cabeca == nil {
		fmt.Println("listaa vazia")
		return
	}

	atual := lista.Cabeca
	for {
		fmt.Printf("%d -> ", atual.Value)
		atual = atual.Next
		if atual == lista.Cabeca {
			break
		}
	}
	fmt.Println()
}

func (lista *listaCircular) InserirNoInicio(value int) {
	novoNo := &No{Value: value}

	if lista.Cabeca == nil {
		novoNo.Next = novoNo
		lista.Cabeca = novoNo
	} else {
		novoNo.Next = lista.Cabeca
		atual := lista.Cabeca
		for atual.Next != lista.Cabeca {
			atual = atual.Next
		}
		atual.Next = novoNo
		lista.Cabeca = novoNo
	}
}

func (lista *listaCircular) RemoverPrimeiro() {
	if lista.Cabeca == nil {
		fmt.Println("A listaa está vazia, nada para excluir.")
		return
	}

	if lista.Cabeca.Next == lista.Cabeca {
		lista.Cabeca = nil
	} else {
		atual := lista.Cabeca
		for atual.Next != lista.Cabeca {
			atual = atual.Next
		}
		atual.Next = lista.Cabeca.Next
		lista.Cabeca = lista.Cabeca.Next
	}
}

func main() {
	lista := &listaCircular{}

	lista.InserirNoInicio(3)
	lista.InserirNoInicio(2)
	lista.InserirNoInicio(1)

	fmt.Println("listaa Circular:")
	lista.Exibir()

	lista.RemoverPrimeiro()
	fmt.Println("Após excluir o primeiro nó:")
	lista.Exibir()
}
