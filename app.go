//THE APP:CONNECTING TO THE DB AND CONNECTINT TO THE ROUTER

package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//establishing a connection with the database and initializing the router
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
}
func (a *App) Run(addr string) {}
