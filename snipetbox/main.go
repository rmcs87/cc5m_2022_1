package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request){
  rw.Write([]byte("Bem vindo ao SnipetBox"))
}

func main() {
	mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  log.Println("Inicializando o servidor na porta: 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
