# Lista Circular

Uma lista circular é uma estrutura de dados na qual cada nó tem um ponteiro para o próximo nó, e o último nó aponta de volta para o primeiro nó, formando um ciclo.

## Exibição dos Nós em uma Lista Circular

```
func ExibirListaCircular(Lista)
    Se Lista está vazia
        Retorne
    Fim Se
    
    NóAtual := Lista.NóInicial
    para
        Escrever(NóAtual.Valor)
        NóAtual = NóAtual.Próximo
    enquanto NóAtual == Lista.NóInicial
Fim Procedimento
```

## Inserção de um Nó no Início da Lista

```
func InserirNoInicio(Lista, Valor)
    NovoNó := CriarNó(Valor)
    
    Se Lista está vazia
        Lista.NóInicial = NovoNó
        NovoNó.Próximo = NovoNó
    Senão
        NovoNó.Próximo = Lista.NóInicial
        Lista.NóInicial = NovoNó
    Fim Se
Fim Procedimento
```

## Exclusão do Primeiro Nó da Lista

```
func ExcluirPrimeiroNó(Lista)
    Se Lista está vazia
        Retorne
    Fim Se
    
    Se Lista possui apenas um nó
        LiberarMemória(Lista.NóInicial)
        Lista.NóInicial = Nulo
    Senão
        NóParaExcluir := Lista.NóInicial
        Lista.NóInicial = NóParaExcluir.Próximo
        LiberarMemória(NóParaExcluir)
    Fim Se
Fim Procedimento
```
