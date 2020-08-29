//PROVIDES THE INFORMATION NEEDED  FOR THE COBNECTIONS TO BE MADE BY THE APP
package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	a.Run(":8010")
}
