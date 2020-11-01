package service

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

func New(db *sql.DB) *Server {
	s := Server{
		DB: db,
	}
	return &s
}

func (s Server) InitializeRoutes() {
	s.Router.HandleFunc("/products", GetProducts).Methods("GET")
	s.Router.HandleFunc("/product", CreateProduct).Methods("POST")
	s.Router.HandleFunc("/product/{id:[0-9]+}", GetProduct).Methods("GET")
	s.Router.HandleFunc("/product/{id:[0-9]+}", UpdateProduct).Methods("PUT")
	s.Router.HandleFunc("/product/{id:[0-9]+}", DeleteProduct).Methods("DELETE")
}
