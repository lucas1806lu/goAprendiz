package main

import (
	"fmt"
	"introducao-testes/enderecos"

)

func main() {
	TipoEndereco := enderecos.TipoDeEndereco("avenida das Rosas")
	fmt.Println(TipoEndereco)
}
