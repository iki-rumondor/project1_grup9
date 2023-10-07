package main

import (
	"log"

	"github.com/iki-rumondor/project1_grup9/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/project1_grup9/internal/adapter/http"
	"github.com/iki-rumondor/project1_grup9/internal/application"
	"github.com/iki-rumondor/project1_grup9/internal/domain"
	"github.com/iki-rumondor/project1_grup9/internal/repository"
	"github.com/iki-rumondor/project1_grup9/internal/routes"
)

func main() {
	db, err := database.NewPostgresDB()
	if err!= nil{
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(domain.Task{})

	taskRepo := repository.NewTaskRepo(db)
	taskService := application.NewTaskService(&taskRepo)
	taskHandler := customHTTP.NewTaskHandler(taskService)

	var PORT = ":8081"
	routes.StartTodoServer(taskHandler).Run(PORT)
}