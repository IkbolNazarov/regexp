package main

import (
	"log"
	"net"
	"regexp/internal/db"
	"regexp/internal/repository"
	"regexp/internal/server"
	"regexp/internal/services"

	"regexp/config"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	DB := db.GetDbConnection()
	config, _, err := config.GetConfig("host")
	if err != nil {
		log.Println(err)
	}

	newRepository := repository.NewRepository(DB)
	newService := services.NewServices(newRepository)
	newServer := server.NewHandler(route, newService)

	newServer.Init()

	address := net.JoinHostPort(config.LocalHost.Host, config.LocalHost.Port)

	err = route.Run(address)
	if err != nil {
		log.Println(err)
	}
}
