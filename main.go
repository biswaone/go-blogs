package main

import (
	"log"
	"net/http"
	"os"
)

// sudo docker run -p 5432:5432 -e POSTGRES_DB=gonotes -e POSTGRES_USER=gonotes -e POSTGRES_PASSWORD=gnotes --name pg postgres

// type App struct {
// 	db      *pgx.Conn
// 	address string
// }

func setupHandlers(mux *http.ServeMux) {
	// mux.HandleFunc("/healthchecker", HealthCheckHandler)
	mux.HandleFunc("/api/user/register", RegisterUserHandler)
}

func main() {
	db := initDB()
	setupDatabaseTables(db)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/api/signup", handleSignUp)
	// // Start the server on port 8080 using the Gorilla Mux router
	// err := http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close(context.Background())
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	mux := http.NewServeMux()
	setupHandlers(mux)
	log.Fatal(http.ListenAndServe(listenAddr, mux))

}
