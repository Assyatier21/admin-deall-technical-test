package main

import (
	"fmt"

	"github.com/assyatier21/admin-deall-technical-test/config"
	cons "github.com/assyatier21/admin-deall-technical-test/models"
	"github.com/assyatier21/admin-deall-technical-test/routes"
)

func main() {
	config.InitDB()
	echo := routes.GetRoutes()
	addres := cons.Addres
	port := cons.Port
	host := fmt.Sprintf("%s:%s", addres, port)
	_ = echo.Start(host)
}
