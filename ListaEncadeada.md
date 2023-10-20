# Lista Duplamente Encadeada Ordenada

Uma lista duplamente encadeada é uma estrutura de dados na qual cada nó possui uma referência tanto para o nó anterior quanto para o nó seguinte na lista. Nesta implementação, os nós estão organizados de forma ordenada. Vamos descrever os algoritmos para as operações comuns em uma lista duplamente encadeada ordenada.

## Definição da Estrutura do Nó

Cada nó na lista duplamente encadeada ordenada terá dois campos: valor e ponteiros para o nó anterior e o nó seguinte.

### Pseudocodigo:
```
Estrutura No:
    Valor
    Ponteiro para o No Anterior
    Ponteiro para o No Seguinte
```
```
InicializarListaDuplamenteEncadeada():
    Inicializar Nó Inicial
    Inicializar Nó Final
    Conectar Nó Inicial com Nó Final
```
```
InserirNo(valor):
    NovoNo <- Criar um novo nó com o valor
    NoAtual <- Nó Inicial
    
    Enquanto NoAtual->Prox não for nulo e NoAtual->Prox->Valor < valor:
        NoAtual <- NoAtual->Prox
    
    NovoNo->Anterior <- NoAtual
    NovoNo->Prox <- NoAtual->Prox
    NoAtual->Prox <- NovoNo
    NovoNo->Prox->Anterior <- NovoNo
```
```
BuscarNo(valor):
    NoAtual <- Nó Sentinela Inicial->Prox
    
    Enquanto NoAtual não for nulo:
        Se NoAtual->Valor igual a valor:
            Retornar NoAtual
        NoAtual <- NoAtual->Prox
    
    Retornar Nulo  # O nó com o valor não foi encontrado
```
```
RemoverNo(valor):
    NoParaRemover <- BuscarNodo(valor)
    
    Se NoParaRemover não for nulo:
        NoParaRemover->Anterior->Prox <- NoParaRemover->Prox
        NoParaRemover->Prox->Anterior <- NoParaRemover->Anterior
```
```
ExibirNo():
    NoAtual <- No Sentinela Inicial->Prox
    
    Enquanto NoAtual não for o No Sentinela Final:
        Exibir NoAtual->Valor
        NoAtual <- NoAtual->Prox
```