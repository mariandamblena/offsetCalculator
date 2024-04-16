package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
)

var (
	server   = flag.String("server", "127.0.0.1", "the database server")
	port     = flag.Int("port", 1433, "the database port")
	database = flag.String("database", "sensors", "the database name")
	/*
		No se usan porque ingresamos con la autentifiacion de windows
		////////////////////////////////////////////////////////////////
		user     = flag.String("user", "", "the database user")
		password = flag.String("password", "", "the database password")

	*/
)

func parseFlags() {
	flag.Parse()
}

func connectionString() string {
	return fmt.Sprintf("server=%s;database=%s;integrated security=true;port=%d", *server, *database, *port)
}

func queryTipoSensor(connectionString, query string) ([]int, []string, error) {
	// Conectar a la base de datos SQL Server
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	// Ejecutar la consulta SQL
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Almacenar los resultados en slices
	var ids []int
	var descripciones []string

	for rows.Next() {
		var id int
		var descripcion string
		if err := rows.Scan(&id, &descripcion); err != nil {
			return nil, nil, err
		}
		ids = append(ids, id)
		descripciones = append(descripciones, descripcion)
	}

	// Verificar errores de iteración
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return ids, descripciones, nil
}

func printResultsTipoSensor(ids []int, descripciones []string) {
	// Imprimir los resultados
	for i := 0; i < len(ids); i++ {
		log.Printf("ID: %d, Descripcion: %s\n", ids[i], descripciones[i])
	}
}

func registerNewSensor(connectionString string) error {
	var id int
	var descripcion string

	fmt.Println("Enter sensor ID:")
	_, err := fmt.Scanln(&id)
	if err != nil {
		return err
	}

	fmt.Println("Enter sensor description:")
	_, err = fmt.Scanln(&descripcion)
	if err != nil {
		return err
	}

	// Consulta SQL para insertar el nuevo sensor
	query := fmt.Sprintf("INSERT INTO TipoSensor (id, Descripcion) VALUES (%d, '%s')", id, descripcion)

	// Conectar a la base de datos SQL Server
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	// Ejecutar la consulta SQL
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Sensor registered successfully")
	return nil
}

/*
Se usa en caso de autentificarse con usuario
///////////////////////////////////////////////////////////////////////////////////////////////////////////

func connectionString() string {
	return fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", *server, *database, *user, *password, *port)
}
*/
