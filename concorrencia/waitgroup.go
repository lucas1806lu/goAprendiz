package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
  
	var waitGroup sync.WaitGroup

waitGroup.Add(4)

go func() {
	escrever("Não para mais")
     waitGroup.Done()
}()
go func() {
	escrever("mais vou executar esse")
     waitGroup.Done()
}()
go func() {
	escrever("Não para mais")
     waitGroup.Done()
}()
go func() {
	escrever("mais vou executar esse")
     waitGroup.Done()
}()
go func() {
	escrever("Não para mais2")
     waitGroup.Done()
}()
go func() {
	escrever("mais vou executar esse2")
     waitGroup.Done()
}()
	
	waitGroup.Wait()

}

func escrever (texto string){

for i := 0;i<=5; i++{
	fmt.Println(texto)
	time.Sleep(time.Second)
}

}