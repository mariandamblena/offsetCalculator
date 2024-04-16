package main

/*
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	parseFlags()

	conn, err := sql.Open("mssql", connectionString())
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Println("Database conection stablished")
	defer conn.Close()

	// Escribir la consulta SQL para insertar valores en la tabla alumno
	insertQuery := "INSERT INTO alumno (legajo, NyA, FechaIng, FechaNac, Mail) VALUES (?, ?, ?, ?, ?)"

	// Preparar la consulta de inserción
	insertStmt, err := conn.Prepare(insertQuery)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer insertStmt.Close()

	// Ejecutar la consulta de inserción con valores de ejemplo
	legajo := 12345
	nya := "Juan Pérez"
	fechaIng := time.Now().Format("2006-01-02") // Formato YYYY-MM-DD
	fechaNac := "1990-01-01"                    // Formato YYYY-MM-DD
	mail := "juanperez@example.com"

	_, err = insertStmt.Exec(legajo, nya, fechaIng, fechaNac, mail)
	if err != nil {
		log.Fatal("Insert failed:", err.Error())
	}

	fmt.Println("Valores insertados correctamente en la tabla alumno.")

	fmt.Printf("bye\n")
}
*/
