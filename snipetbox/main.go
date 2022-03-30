package main
import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    http.NotFound(rw, r)
    return
  } 
  rw.Write([]byte("Bem vindo ao SnipetBox"))
}
func showSnippet(rw http.ResponseWriter, r *http.Request){
  rw.Write([]byte("Mostrar um Snippet especifico"))
}
func createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar novo snippet"))
}
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
