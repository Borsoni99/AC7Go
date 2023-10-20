package main

import "fmt"


type No struct {
	Value int
	Proximo  *No
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
		atual = atual.Proximo
		if atual == lista.Cabeca {
			break
		}
	}
	fmt.Println()
}

func (lista *listaCircular) InserirNoInicio(value int) {
	novoNo := &No{Value: value}

	if lista.Cabeca == nil {
		novoNo.Proximo = novoNo
		lista.Cabeca = novoNo
	} else {
		novoNo.Proximo = lista.Cabeca
		atual := lista.Cabeca
		for atual.Proximo != lista.Cabeca {
			atual = atual.Proximo
		}
		atual.Proximo = novoNo
		lista.Cabeca = novoNo
	}
}

func (lista *listaCircular) RemoverPrimeiro() {
	if lista.Cabeca == nil {
		fmt.Println("A listaa está vazia, nada para excluir.")
		return
	}

	if lista.Cabeca.Proximo == lista.Cabeca {
		lista.Cabeca = nil
	} else {
		atual := lista.Cabeca
		for atual.Proximo != lista.Cabeca {
			atual = atual.Proximo
		}
		atual.Proximo = lista.Cabeca.Proximo
		lista.Cabeca = lista.Cabeca.Proximo
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
