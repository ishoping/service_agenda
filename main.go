package main

import (
	"os"

	"service_agenda/entity"
	"service_agenda/service"

	flag "github.com/spf13/pflag"
)

const (
	//PORT .
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	dbPath := flag.StringP("dbPath", "d", "", "path for database")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	if len(*dbPath) == 0 {
		os.Mkdir("data", 0755)
		*dbPath = "data/test.db"
	}

	entity.Connectdb(*dbPath)
	server := service.NewServer()
	server.Run(":" + port)
}
