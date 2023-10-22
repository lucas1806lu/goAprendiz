package main

import (
	"fmt"
)

func main() {

	// o "3" signifa o buffer quantidade de canais qe deve esperar os canais para executar o resto do programa  
	canal := make(chan string, 3)

	canal <- "ola mundo"
	canal <- "Programando em Go "
	canal <- "Progamei go new app"

	mensagem := <- canal
	mensagem2 := <- canal
	mensagem3 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
	

}
