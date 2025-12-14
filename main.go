package main

import (
	"context"
	"log"
	"project-app-inventaris-cli-azwin/cmd"
	"project-app-inventaris-cli-azwin/database"
	"project-app-inventaris-cli-azwin/handler"
	"project-app-inventaris-cli-azwin/repository"
	"project-app-inventaris-cli-azwin/service"

	_ "github.com/lib/pq"
)

func main() {
	//init DB connection
	db, err := database.InitDB()

	//check database connection
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}else{
		log.Println("Database connection successful")
	}
	
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}else{
		log.Println("Database ping successful")
	}

	defer db.Close(context.Background())

	//init object
	repoCategory := repository.NewrepoCategory(db)
	serviceCategory := service.NewServiceCategory(&repoCategory)
	handlerCategory := handler.NewHandlerCategory(&serviceCategory)

	repoManagement := repository.NewrepoManagement(db)
	serviceManagement := service.NewServiceManagement(&repoManagement)
	handlerManagement := handler.NewHandlerManagement(&serviceManagement)

	repoOld := repository.NewRepositoryOld(db)
	serviceOld := service.NewServiceOld(&repoOld)
	handlerOld := handler.NewHandlerOld(&serviceOld)

	repoReport := repository.NewRepositoryReport(db)
	serviceReport := service.NewServiceReport(&repoReport)
	handlerReport := handler.NewHandlerReport(&serviceReport)

	// Run in interactive mode
	cmd.Home(handlerCategory, handlerManagement, handlerOld, handlerReport)
}
