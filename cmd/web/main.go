package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	adrr := flag.String("addr", ":4000", "network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	mux := app.routes()
	infoLog.Printf("starting server at %v\n", *adrr)
	srv := &http.Server{Handler: mux, ErrorLog: errorLog, Addr: *adrr}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
