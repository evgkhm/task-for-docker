package main

import (
	"net/http"
	"task-for-docker/internal/config"
	"task-for-docker/internal/controller/handlers"
	"task-for-docker/internal/repository/postgres"
	"task-for-docker/internal/usecase"
)

func init() {
	config.InitAll([]config.Config{
		handlers.HTTPConfig{},
		postgres.Conf{},
	})
}

func main() {
	postgresDB, err := postgres.NewPostgresDB(postgres.Config)
	if err != nil {
		panic(err)
	}

	postgresRepository := postgres.New(postgresDB)

	useCases := usecase.New(postgresRepository)

	APIHandlers := handlers.New(useCases)
	r := APIHandlers.InitRoutes()

	http.ListenAndServe(":"+handlers.HTTP.Port, r)

}
