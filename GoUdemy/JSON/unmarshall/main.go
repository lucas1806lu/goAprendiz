package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {

	//para ignorar um campo do Json "-" ele não vai reconhecer a identificação
	Nome string `json: "-"`   //"nome"

	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {

    //para ignorar um campo do Json "-" ele não vai reconhecer a identificação
	cachorroEmJSON := `{"-" : "Rex", "raca" : "boxer", "idade" : 5}`

	var c cachorro
	// converte um JSON em struc ou slice
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome" : "Apolo", "raca" : "labrador", "idade" : "9"}`

	c2 := make(map[string]string)
	// converte um JSON em map
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
