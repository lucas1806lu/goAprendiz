package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


type cachorro struct{
	Nome string  `json: "nome"`
	Raca string  `json: "raça"`
	Idade uint   `json: "idade"`
}

func main(){

	 c := cachorro {"Rex", "Dalmata", 5}
	 fmt.Println(c)

	 // transforma dados no formato JSON
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)
	//converte dados recebido em bytes, em estrtura de leitura
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))


	c2 := map[string]string{
		"nome" : "thor",
		"raça" : "pitbull",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)

	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))


}