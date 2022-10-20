package main

import (
	"fmt"
	"log"
	httpPkg "net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
	"github.com/seniorescobar/alchemy-test/internal/gateway/http"
	"github.com/seniorescobar/alchemy-test/internal/storage/mysql"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	repo := mysql.NewSpacecraftRepository()
	svc := spacecraft.NewService(repo)
	gw := http.NewSpacecraftGateway(svc)

	if err := startServer(gw); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}

func startServer(gw *http.SpacecraftGateway) error {
	router := mux.NewRouter()
	router.HandleFunc("/", gw.List).Methods(httpPkg.MethodGet)
	router.HandleFunc("/{id}", gw.Get).Methods(httpPkg.MethodGet)

	srv := &httpPkg.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
