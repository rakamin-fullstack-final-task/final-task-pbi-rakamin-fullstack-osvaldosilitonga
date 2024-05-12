package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cant load .env file")
	}
}

func main() {
	r := gin.Default()

	router.Route(r)

	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
