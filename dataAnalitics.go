package main

import (
	"database/sql"
	"fmt"
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

// convertExcelDataToDataset convierte los datos del archivo Excel al formato de datos requerido
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

			// Convertir la cadena de fecha y hora a un objeto time.Time
			timestamp, err := time.Parse("2/1/06 15:04", timestampStr)
			if err != nil {
				log.Printf("Error parsing timestamp at row %d: %v", rowIndex+1, err)
				continue
			}

			// Convertir las cadenas de humedad y temperatura a flotantes
			hum, err := strconv.ParseFloat(strings.Replace(humStr, ",", ".", -1), 64)
			if err != nil {
				log.Printf("Error parsing humidity at row %d: %v", rowIndex+1, err)
				continue
			}
			temp, err := strconv.ParseFloat(strings.Replace(tempStr, ",", ".", -1), 64)
			if err != nil {
				log.Printf("Error parsing temperature at row %d: %v", rowIndex+1, err)
				continue
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
