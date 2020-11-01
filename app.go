package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/DavidRS91/simple-go-server/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// var createTableQuery string = "CREATE TABLE IF NOT EXISTS products( id SERIAL, name TEXT NOT NULL, price NUMERIC(10,2) NOT NULL DEFAULT 0.00, CONSTRAINT products_pkey PRIMARY KEY (id));"

func (a *App) Initialize(user, password, dbname, host, port, sslmode string) {
	fmt.Println("Initializing...")
	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, password, dbname, host, port, sslmode)
	var err error
	fmt.Println("Connecting to db...")
	a.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running migrations...")
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		sslmode,
	)
	m, err := migrate.New("file://data/migrations", postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	server := service.New(a.DB, a.Router)
	server.InitializeRoutes()
}

func (a *App) Run(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	fmt.Printf("Listening on %s", addr)
	handler := c.Handler(a.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
