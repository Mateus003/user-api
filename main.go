package main

import (
	"log"

	"github.com/Mateus003/user-api/src/configuration/logger"
	"github.com/Mateus003/user-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
  logger.Info("About start user application")
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  router := gin.Default()

  routes.InitRoutes(&router.RouterGroup)

  if err := router.Run(":8080") ; err!= nil{
	log.Fatal(err)
  }
}