package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
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

/*
Se usa en caso de autentificarse con usuario
///////////////////////////////////////////////////////////////////////////////////////////////////////////

func connectionString() string {
	return fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", *server, *database, *user, *password, *port)
}
*/

func parseFlags() {
	flag.Parse()
}

func connectionString() string {
	return fmt.Sprintf("server=%s;database=%s;integrated security=true;port=%d", *server, *database, *port)
}

func queryTipoSensor(connectionString string) ([]int, []string, error) {
	// Cadena de consulta SQL
	query := "SELECT id, Descripcion FROM TipoSensor"

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

func querySensorView(connectionString string) ([]int, []string, error) {
	// Cadena de consulta SQL
	query := "SELECT * FROM SensorView"

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

func registerNewSensorType(connectionString string) error {
	var id int
	var descripcion string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter sensor ID:")
	_, err := fmt.Fscanln(reader, &id)
	if err != nil {
		return err
	}

	fmt.Println("Enter sensor description:")
	_, err = fmt.Fscanln(reader, &descripcion)
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

func uploadDatatoSensor(connectionString string) error {
	var id int
	var descripcion string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select sensor ID:")
	_, err := fmt.Fscanln(reader, &id)
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

func registerNewSensor(connectionString string) error {
	var SN string
	var sensorTypeID int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter sensor Serial number:")
	_, err := fmt.Fscanln(reader, &SN)
	if err != nil {
		return err
	}

	fmt.Println("We have the following types of sensors:")
	ids, descriptions, errQuery := queryTipoSensor(connectionString)
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	printResultsTipoSensor(ids, descriptions)
	fmt.Println("")

	fmt.Println("Enter sensor type ID:")
	_, err = fmt.Fscanln(reader, &sensorTypeID)
	if err != nil {
		return err
	}

	// SQL query to insert the new sensor
	registerQuery := fmt.Sprintf("INSERT INTO Sensor (SerialNumber, TipoSensorId) VALUES ('%s', %d)", SN, sensorTypeID)

	// Connect to the SQL Server database
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	// Execute the SQL query
	_, err = db.Exec(registerQuery)
	if err != nil {
		return err
	}

	fmt.Println("Sensor registered successfully")
	return nil
}
