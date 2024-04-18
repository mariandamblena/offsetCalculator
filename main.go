package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
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

			fmt.Println("What action do you want to perform? \n [1] Register new sensor in Stock \n [2] Add a new Sensor Type \n [3] Sensor list \n [4] Delete sensor \n [5] Delete sensor type \n ")

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
				fmt.Println("Insert the serial number to be deleted...")
				// Prompt for serial number input and delete the sensor
				err := deleteSensorBySerialNumber(connectionString())
				if err != nil {
					log.Fatal(err)
				}

			case 5:
				fmt.Print("\033[H\033[2J")
			default:
				fmt.Println("Invalid option selected.")
				fmt.Print("\033[H\033[2J")
			}

		case 2:
			fmt.Println("Upload sensor datalogs...")
			// Implement uploading sensor datalogs functionality.

		case 3:
			fmt.Println("Calculating Offset...")
			// Call a function to handle calculating offset.

		default:
			fmt.Println("Invalid option selected.")
		}
	}
}
