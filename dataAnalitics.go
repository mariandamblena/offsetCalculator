package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

type Data struct {
	Timestamp time.Time
	Value1    float64 // Temperatura
	Value2    float64 // Humedad
}

func insertDataIntoDatalog(connectionString string, dataset []Data, serialNumber string) error {
	// Connect to the SQL Server database
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL statement for inserting data into Datalog
	query := "INSERT INTO Datalog (Timestamp, Value1, Value2, SensorSerialNumber) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Iterate through the dataset and insert each data point into Datalog
	for _, data := range dataset {
		_, err := stmt.Exec(data.Timestamp, data.Value1, data.Value2, serialNumber)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Data inserted into Datalog for sensor with Serial Number %s successfully\n", serialNumber)
	return nil
}

func convertExcelDataToDataset(file *xlsx.File) []Data {
	var dataset []Data

	// Iterar sobre las hojas del archivo Excel
	for _, sheet := range file.Sheets {
		// Iterar sobre las filas de la hoja
		for rowIndex, row := range sheet.Rows {
			// Omitir la primera fila (encabezados)
			if rowIndex == 0 {
				continue
			}

			// Leer los valores de la fila
			timestampStr := row.Cells[1].String() // Fecha-Hora
			humStr := row.Cells[2].String()       // Humedad
			tempStr := row.Cells[3].String()      // Temperatura

			// Convertir el timestamp de Excel a time.Time
			timestampDecimal, err := strconv.ParseFloat(timestampStr, 64)
			if err != nil {
				log.Printf("Error parsing timestamp at row %d: %v", rowIndex+1, err)
				continue
			}
			timestamp, err := convertDecimalToTime(timestampDecimal)
			if err != nil {
				log.Printf("Error converting timestamp at row %d: %v", rowIndex+1, err)
				continue
			}

			// Manejar los valores de humedad y temperatura no válidos
			var hum, temp float64
			if humStr == "----" || tempStr == "----" {
				// Si los valores son '----', establecer a un valor predeterminado o manejar según sea necesario
				hum = 0.0  // Valor predeterminado para humedad
				temp = 0.0 // Valor predeterminado para temperatura
			} else {
				// Convertir las cadenas de humedad y temperatura a flotantes
				var err error
				hum, err = strconv.ParseFloat(strings.Replace(humStr, ",", ".", -1), 64)
				if err != nil {
					log.Printf("Error parsing humidity at row %d: %v", rowIndex+1, err)
					continue
				}
				temp, err = strconv.ParseFloat(strings.Replace(tempStr, ",", ".", -1), 64)
				if err != nil {
					log.Printf("Error parsing temperature at row %d: %v", rowIndex+1, err)
					continue
				}
			}

			// Crear un nuevo punto de datos y agregarlo al conjunto de datos
			data := Data{
				Timestamp: timestamp,
				Value1:    temp,
				Value2:    hum,
			}
			dataset = append(dataset, data)
		}
	}

	return dataset
}

func convertDecimalToTime(decimal float64) (time.Time, error) {
	// Extraer la parte entera y decimal
	days := int(decimal)
	fractionalPart := decimal - float64(days)

	// Convertir la parte decimal a horas, minutos y segundos
	secondsInDay := int(fractionalPart * 86400) // 86400 segundos en un día
	hours := secondsInDay / 3600
	secondsInDay %= 3600
	minutes := secondsInDay / 60
	seconds := secondsInDay % 60

	// Construir la fecha y hora
	referenceDate := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC) // Fecha de referencia de Excel
	targetDate := referenceDate.AddDate(0, 0, days).Add(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second)

	return targetDate, nil
}

func listFilesInFolder(folderPath string) ([]string, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func selectFileByIndex(fileNames []string) (string, error) {
	fmt.Println("Select a file by entering its index:")
	for i, name := range fileNames {
		fmt.Printf("[%d] %s\n", i+1, name)
	}

	var selectedIndex int
	fmt.Print("Enter the index of the file: ")
	_, err := fmt.Scanln(&selectedIndex)
	if err != nil {
		return "", err
	}

	if selectedIndex < 1 || selectedIndex > len(fileNames) {
		return "", fmt.Errorf("invalid index")
	}

	selectedFileName := fileNames[selectedIndex-1]
	return selectedFileName, nil
}
