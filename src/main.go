package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"go-clean-arch/src/config"
	"go-clean-arch/src/middleware"

	_entity "go-clean-arch/src/api/entity"
	_handler "go-clean-arch/src/api/handler"
	_repo "go-clean-arch/src/api/repo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
			log.Fatal("Error getting env")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	timeout := os.Getenv("TIMEOUT")
	if timeout == "" {
		timeout = "2"
	}

	i, _ := strconv.Atoi(timeout)
	timeoutContext := time.Duration(i) * time.Second

	config.InitSql()
	db := config.GetSqlDB()

	repo := _repo.InitRepo(db)
	entity := _entity.InitEntity(repo, timeoutContext)

	r := gin.Default()
	r.Use(middleware.CORS())
	api := r.Group("/")

	_handler.InitHandler(api, entity)
	r.Run(":" + port)

}
