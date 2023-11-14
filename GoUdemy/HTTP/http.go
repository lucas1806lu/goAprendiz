package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ola mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("carregar pagina de usuarios!"))
}

func main() {
	//HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
	// CLIENTE-SERVIDOR
	// REQUEST - RESPONSE      -
	//ROTAS
	//URI - identificador de recursos
	// Metodos - GET, POST, PUT, DELETE

	// esse codigo cria a rota recebendo 2 parametros 1º a URI e 2º a função que sera executada
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)



	//com essa linha de codigo a consegue counicar com a porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))

}
