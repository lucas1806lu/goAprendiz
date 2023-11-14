package main

import "fmt"


func fibonacci (posicao int) int{
 if posicao <= 1{
	return posicao
 }
  return fibonacci (posicao - 2) + fibonacci (posicao -1)
}

func worker(tarefas <-chan int, resultados chan<- int){
for numeros := range tarefas{
	resultados <- fibonacci(numeros) 
}

}

func main(){

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
//no Padrao worker Pools nesse exemplos temos 4 tafas executando os resultados
//acelerando a resposta na tela  
	go worker (tarefas, resultados)
	go worker (tarefas, resultados)
	go worker (tarefas, resultados)
	go worker (tarefas, resultados)
	
	for i:=0; i<45; i++{
		tarefas <- i
	}

	for i:=0; i<45; i++{
		resultado :=  <-resultados
		fmt.Println(resultado)
	}

}