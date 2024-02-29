package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/claudineyveloso/api-zssn/internal/db"
	"github.com/claudineyveloso/api-zssn/routes"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user_api_zssn"
	password = "pwd_api_zssn"
	dbname   = "api_zssn"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	fmt.Println("Conectou com o banco de dados!")


	dbUser := db.New(dbConn)

	http.HandleFunc("/create_user", func(w http.ResponseWriter, r *http.Request) {
		routes.CreateUser(w, r, dbUser)
  })

	http.HandleFunc("/get_users", func(w http.ResponseWriter, r *http.Request) {
    routes.GetUsers(w, r, dbUser)
  })

	http.ListenAndServe(":8080", nil)

}
