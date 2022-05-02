package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

  _ "github.com/go-sql-driver/mysql"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
}

//curl -i -X GET http://localhost:4000/snippet/create
func main() {
	// nome da flag, valor padra e descrição
	addr := flag.String("addr", ":4000", "Porta da Rede")
  dsn := flag.String("dsn",
                     "RNNRfSNISQ:jMFmtb1Hby@tcp(remotemysql.com)/RNNRfSNISQ?parseTime=true", 
                     "MySql DSN")
  
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil{
    errorLog.Fatal(err)
  }
  defer db.Close()
  
  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
  }

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }
  
	infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error){
  db, err := sql.Open("mysql", dsn)
  if err!= nil{
    return nil, err
  }
  if err = db.Ping(); err != nil{
    return nil, err
  }
  return db, nil
}