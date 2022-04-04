package main

import (
	"log"
	"net/http"
)

//curl -i -X GET http://localhost:4000/snippet/create
func main() {
	mux := http.NewServeMux()
  
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)
  
  log.Println("Inicializando o servidor na porta: 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
