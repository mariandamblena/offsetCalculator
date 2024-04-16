package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var selector int

func main() {

	conn, err := sql.Open("mssql", connectionString())
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	fmt.Println("Database connection established")

	fmt.Println("Welcome to the Offset Calculator, what action you want to perform? \n [1] Register a new sensor \n [2] Upload sensor datalogs \n [3] Calculate Offset \n [4] Run a Modbus datalogging (coming soon)")

	_, scanErr := fmt.Scanln(&selector)
	if scanErr != nil {
		log.Fatal("Error reading input:", scanErr)
	}

	switch selector {
	case 1:
		fmt.Println("Registering a new sensor...")
		// Llamar a la función queryTipoSensor para saber que sensores estan registrados

		query := "SELECT id, Descripcion FROM TipoSensor"

		ids, descripciones, err := queryTipoSensor(connectionString(), query)
		if err != nil {
			log.Fatal(err)
		}

		// Imprimir los resultados

		printResultsTipoSensor(ids, descripciones)
		// if is required, call a function to ask for serial number and sersor description to register en sql database

	case 2:
		fmt.Println("Uploading sensor datalogs...")
		// Call a function to handle uploading sensor datalogs.
	case 3:
		fmt.Println("Calculating Offset...")
		// Call a function to handle calculating offset.
	case 4:
		fmt.Println("Running Modbus datalogging (coming soon)...")
		// Provide a message indicating that this feature is not yet implemented.
	default:
		fmt.Println("Invalid option selected.")
	}
}
