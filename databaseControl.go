package main

import (
	"database/sql"
	"flag"
	"fmt"
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

func queries(connectionString, query string) ([]int, []string, error) {
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

/*
Se usa en caso de autentificarse con usuario
///////////////////////////////////////////////////////////////////////////////////////////////////////////

func connectionString() string {
	return fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", *server, *database, *user, *password, *port)
}
*/
