package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/tealeg/xlsx"
)

var selector int8

func main() {
	fmt.Print("\033[H\033[2J")

	for {
		conn, err := sql.Open("mssql", connectionString())
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		defer conn.Close()

		fmt.Println("Welcome to the Offset Calculator, what action do you want to perform? \n [1] Sensors management \n [2] Upload sensor datalogs \n [3] Calculate Offset \n ")

		_, scanErr := fmt.Scanln(&selector)
		if scanErr != nil {
			log.Fatal("Error reading input:", scanErr)
		}

		switch selector {
		case 1:
			fmt.Print("\033[H\033[2J")

			fmt.Println("What action do you want to perform? \n [1] Register new sensor in Stock \n [2] Add a new Sensor Type \n [3] Sensor list \n [4] Sensor type list \n [5] Delete sensor \n [6] Delete sensor type \n ")

			var case1Selector int8

			_, scanErr := fmt.Scanln(&case1Selector)
			if scanErr != nil {
				log.Fatal("Error reading input:", scanErr)
			}

			switch case1Selector {
			case 1:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Registering a new sensor...")
				err := registerNewSensor(connectionString())
				if err != nil {
					log.Fatal(err)
				}

			case 2:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Registering a new sensor type...")
				err := registerNewSensorType(connectionString())
				if err != nil {
					log.Fatal(err)
				}

			case 3:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Sensor list...")
				// Retrieve sensor information
				SNs, descripciones, ids, errQuery := querySensorView(connectionString())
				if errQuery != nil {
					log.Fatal(errQuery)
				}

				// Display sensor information
				for i := range SNs {
					fmt.Printf("Serial Number: %s, Sensor Type: %s, Sensor type Id: %d\n", SNs[i], descripciones[i], ids[i])
				}

			case 4:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Sensor type list...")
				// Retrieve sensor information
				fmt.Println("We have the following types of sensors:")
				ids, descriptions, errQuery := queryTipoSensor(connectionString())
				if errQuery != nil {
					log.Fatal(errQuery)
				}
				printResultsTipoSensor(ids, descriptions)
				fmt.Println("")

			case 5:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Insert the serial number to be deleted...")
				// Prompt for serial number input and delete the sensor
				err := deleteSensorBySerialNumber(connectionString())
				if err != nil {
					log.Fatal(err)
				}
			case 6:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Delete sensor type by ID...")
				// Delete a sensor type by ID
				err := deleteTipoSensorByID(connectionString())
				if err != nil {
					log.Fatal(err)
				}

			default:
				fmt.Println("Invalid option selected.")
				fmt.Print("\033[H\033[2J")
			}

		case 2:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Upload sensor datalogs...")
			// Solicitar al usuario que ingrese la ruta del archivo Excel
			var filePath string
			fmt.Print("Enter the path to the Excel file: ")
			fmt.Scanln(&filePath)

			// Leer el archivo Excel
			xlFile, err := xlsx.OpenFile(filePath)
			if err != nil {
				log.Fatal(err)
			}

			// Convertir los datos del archivo Excel al formato de datos requerido
			dataset := convertExcelDataToDataset(xlFile)

			// Imprimir el conjunto de datos convertido
			for _, data := range dataset {
				fmt.Printf("Timestamp: %s, Value1: %.2f, Value2: %.2f\n", data.Timestamp, data.Value1, data.Value2)
			}
		case 3:
			fmt.Println("Calculating Offset...")
			// Call a function to handle calculating offset.

		default:
			fmt.Println("Invalid option selected.")
		}
	}
}
