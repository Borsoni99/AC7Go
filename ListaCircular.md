Exibição dos nós em uma lista circular:
```
def exibir(){
  no_atual = primeiro_no
  faca
      exibir no_atual.valor
      no_atual = no_atual.proximo
  enquanto no_atual ≠ primeiro_no
}
```
Inserção de um nó no início da lista:
```
def inserir(){
  novo_no = criar_novo_no(valor)
  novo_no.proximo = primeiro_no
  ultimo_no = encontrar_ultimo_no()
  ultimo_no.proximo = novo_no
  primeiro_no = novo_no
}
```
Exclusão do primeiro nó da lista:
```
def remover(){
  nó_a_remover = primeiro_nó
  primeiro_nó = nó_a_remover.próximo
  último_nó = encontrar_último_nó()
  último_nó.próximo = primeiro_nó
}
```
