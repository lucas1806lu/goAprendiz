package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil{
		log.Fatal(erro)
	}

	/*  */
	//fmt.Println(db)
	fmt.Println("conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil{
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
