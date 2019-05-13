package app

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type App struct {
// 	Router *mux.Router
// 	DB     *sql.DB
// }

// func (a *App) Initialize(user, password, dbname string) {
// 	connectionString :=
// 		"user=postgres password=admin dbname=brisamax sslmode=disable"
// 		//fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, ssl)

// 	var err error
// 	a.DB, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	a.Router = mux.NewRouter()
// 	a.initializeRoutes()
// }

// // Run starts the app and serves on the specified addr
// func (a *App) Run(addr string) {
// 	fmt.Println("Successfully connected!")
// 	log.Fatal(http.ListenAndServe(addr, a.Router))
// }

// func (a *App) initializeRoutes() {
// 	a.Router.HandleFunc("/programs", a.getPrograms).Methods("GET")
// 	a.Router.HandleFunc("/program", a.createProgram).Methods("POST")
// 	a.Router.HandleFunc("/program/{id:[0-9]+}", a.getProgram).Methods("GET")
// 	a.Router.HandleFunc("/program/{id:[0-9]+}", a.updateProgram).Methods("PUT")
// 	a.Router.HandleFunc("/program/{id:[0-9]+}", a.deleteProgram).Methods("DELETE")
// }
