package main

import (
	"admin/config"
	cons "admin/models"
	"admin/routes"
	"fmt"
)

func main() {
	config.InitDB()
	echo := routes.GetRoutes()
	addres := cons.Addres
	port := cons.Port
	host := fmt.Sprintf("%s:%s", addres, port)
	_ = echo.Start(host)
}
