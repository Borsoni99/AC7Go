package main

import "fmt"

type No struct {
	Valor      int
	Ant   *No
	Prox    *No
}

type ListaDuplamenteEncadeada struct {
	NoInicial *No
	NoFinal   *No
}

func InicializarListaDuplamenteEncadeada() *ListaDuplamenteEncadeada {
	NoInicial := &No{}
	NoFinal := &No{}
	NoInicial.Prox = NoFinal
	NoFinal.Ant = NoInicial

	return &ListaDuplamenteEncadeada{NoInicial: NoInicial, NoFinal: NoFinal}
}

func (lista *ListaDuplamenteEncadeada) InserirNo(valor int) {
	novoNo := &No{Valor: valor}
	NoAtual := lista.NoInicial

	for NoAtual.Prox != lista.NoFinal && NoAtual.Prox.Valor < valor {
		NoAtual = NoAtual.Prox
	}

	novoNo.Ant = NoAtual
	novoNo.Prox = NoAtual.Prox
	NoAtual.Prox.Ant = novoNo
	NoAtual.Prox = novoNo
}

func (lista *ListaDuplamenteEncadeada) BuscarNo(valor int) *No {
	NoAtual := lista.NoInicial.Prox

	for NoAtual != lista.NoFinal {
		if NoAtual.Valor == valor {
			return NoAtual
		}
		NoAtual = NoAtual.Prox
	}
	return nil
}

func (lista *ListaDuplamenteEncadeada) RemoverNo(valor int) {
	NoParaRemover := lista.BuscarNo(valor)

	if NoParaRemover != nil {
		NoParaRemover.Ant.Prox = NoParaRemover.Prox
		NoParaRemover.Prox.Ant = NoParaRemover.Ant
	}
}

func (lista *ListaDuplamenteEncadeada) ExibirNos() {
	NoAtual := lista.NoInicial.Prox

	for NoAtual != lista.NoFinal {
		fmt.Printf("%d -> ", NoAtual.Valor)
		NoAtual = NoAtual.Prox
	}
	fmt.Println()
}
func main() {
	lista := InicializarListaDuplamenteEncadeada()
	lista.InserirNo(5)
	lista.InserirNo(3)
	lista.InserirNo(8)
	lista.InserirNo(1)

	fmt.Println("Nós da lista:")
	lista.ExibirNos()

	valorBuscado := 8
	NoBuscado := lista.BuscarNo(valorBuscado)
	if NoBuscado != nil {
		fmt.Printf("Nó com valor %d encontrado.\n", valorBuscado)
	} else {
		fmt.Printf("Nó com valor %d não encontrado.\n", valorBuscado)
	}

	valorRemover := 3
	lista.RemoverNo(valorRemover)
	fmt.Printf("Removendo nó com valor %d.\n", valorRemover)

	fmt.Println("Nós restantes na lista após a remoção:")
	lista.ExibirNos()
}
