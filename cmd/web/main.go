package main

import (
	"os"
	"log"
	"flag"
	"net/http"
	"database/sql"

	"snipbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn  := flag.String("dsn", "web:pass@/snipbox?parseTime=true", "MariaDB data source name")

	flag.Parse()

	errorLog := log.New(os.Stderr, "\033[31mERROR\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog  := log.New(os.Stdout, "\033[33mINFO\033[0m\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on http://localhost%v", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
