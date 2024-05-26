package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaydto/goApiMyql/service/users"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}

}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := users.NewStore(s.db)

	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	log.Println("listen on ", s.addr)

	return http.ListenAndServe(s.addr, router)

}
