package main

import (
	"flag"
	"fmt"
)

var (
	server = flag.String("server", "127.0.0.1", "the database server")
	port   = flag.Int("port", 1434, "the database port")
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
	return fmt.Sprintf("server=%s;integrated security=true;port=%d", *server, *port)
}

/*
Se usa en caso de autentificarse con usuario
///////////////////////////////////////////////////////////////////////////////////////////////////////////

func connectionString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
}
*/
