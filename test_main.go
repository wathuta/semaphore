//TESTING THE APP
//  *THE DATA BASE
//  *CONNECTION
package main

import (
	"log"
	"os"
	"testing"
)

// this represents the application we want to test
var a App

//ensuring that the database is weell set since we are running the tests through the db
//the db should be cleaned up after all tests have been finished

//TestMain is executed before all other tests are executed
func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id),
)`

//ensureTableExists function is used to ,make sure that the table  required for the testing is available
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
