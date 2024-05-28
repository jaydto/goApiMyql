package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jaydto/goApiMyql/service/product"
	"github.com/jaydto/goApiMyql/service/users"
    httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "github.com/jaydto/goApiMyql/docs"
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

// Run starts the HTTP server and serves the API endpoints.
//
// It uses the Swagger UI to generate the documentation. You can access it by
// navigating to the root path of the API server.
//
// The Swagger UI is authenticated by default. Use the following credentials
// to access it:
//
//     Username: admin
//     Password: password
//
// To disable the authentication, update the main.go file and set the
// `disableAuth` variable to true before running the server.
//
// To generate the Swagger documentation, you can use the following command:
//
//     go get github.com/swaggo/swag/cmd/swag
//     swag init
//
// This will generate the swagger.json file in the root directory of the project.
//
// You can configure the username and password for the Swagger UI by updating
// the following environment variables:
//
//     SWAGGER_USERNAME=admin
//     SWAGGER_PASSWORD=password
//
// If you want to disable the authentication, set the following environment
// variable to true:
//
//     SWAGGER_DISABLE_AUTH=true


func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := users.NewStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)


	// Serve the Swagger documentation
    swaggerURL := "./docs"
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    // Serve the Swagger JSON file
    router.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, swaggerURL)
    })


	log.Println("listen on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
