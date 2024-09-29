package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BrancheDeboua/url-shortener/internal/controller"
	"github.com/BrancheDeboua/url-shortener/internal/database"
)

func Serve(port string) {
	connStr := os.Getenv("DATABASE_URL")
	databaseConnector := database.NewPostgresConnector(connStr)

	shortener := controller.NewShortener(databaseConnector)

	router := router(shortener)

	fmt.Println("Listening on port ", port)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println("Error while starting server: ", err)
	}
}
