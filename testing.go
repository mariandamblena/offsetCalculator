package main

import "log"

func test() {
	// Cadena de conexión a la base de datos SQL Server
	connectionString := connectionString()

	// Consulta SQL para seleccionar todos los registros de la tabla TipoSensor
	query := "SELECT id, Descripcion FROM TipoSensor"

	// Llamar a la función queries
	ids, descripciones, err := queries(connectionString, query)
	if err != nil {
		log.Fatal(err)
	}

	// Imprimir los resultados
	for i := 0; i < len(ids); i++ {
		log.Printf("ID: %d, Descripcion: %s\n", ids[i], descripciones[i])
	}
}
