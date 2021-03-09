package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var arguments = os.Args

func dsn (dbName string) string {
	mysqlUser := arguments[1]
	mysqlPass := arguments[2]
	mysqlHost := arguments[3]
	mysqlDB	  := arguments[4]

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlDB)
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select 1")
	
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, "You have an error in your SQL syntax\n")
		return
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Println("Success")
	fmt.Fprintf(w, "ProxySql OK\n")
}

func main() {
	fmt.Printf("Starting server at port 8080\n")
    http.HandleFunc("/healtcheck", sqlQuery)
    http.ListenAndServe(":8080", nil)
}