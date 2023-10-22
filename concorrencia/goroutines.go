package main

import (
	"fmt"
	"time"
)

func main(){
  
	// gorotines usa "go" na frente da execção para execultar a proxima chamada
	go escrever("Não para mais") // gorountines
	escrever("mais vou executar esse")

}

func escrever (texto string){

for{
	fmt.Println(texto)
	time.Sleep(time.Second)
}

}