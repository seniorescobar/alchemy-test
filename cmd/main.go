package main

import (
	"database/sql"
	"fmt"
	"log"
	httpPkg "net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlPkg "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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
	db, err := initDb()
	if err != nil {
		return fmt.Errorf("error initalizing db conn: %w", err)
	}
	defer db.Close()

	repo := mysql.NewSpacecraftRepository(db)
	svc := spacecraft.NewService(repo)
	gw := http.NewSpacecraftGateway(svc)

	if err := startServer(gw); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}

func initDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "alchemy-test:qwerty123@tcp(localhost:3306)/alchemy-test-db?multiStatements=true")
	if err != nil {
		return nil, err
	}

	driver, err := mysqlPkg.WithInstance(db, &mysqlPkg.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://../internal/storage/mysql/migrations", "mysql", driver)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			return nil, err
		}
	}

	return db, nil
}

func startServer(gw *http.SpacecraftGateway) error {
	router := mux.NewRouter()
	router.HandleFunc("/", gw.List).Methods(httpPkg.MethodGet)
	router.HandleFunc("/{id}", gw.Get).Methods(httpPkg.MethodGet)
	router.HandleFunc("/", gw.Create).Methods(httpPkg.MethodPut)
	router.HandleFunc("/{id}", gw.Update).Methods(httpPkg.MethodPatch)
	router.HandleFunc("/{id}", gw.Delete).Methods(httpPkg.MethodDelete)

	srv := &httpPkg.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
