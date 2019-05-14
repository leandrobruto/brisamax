package genero

import (
	"brisamax/app/models/genero"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		"user=postgres password=root dbname=BRISAMAX sslmode=disable"

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("Successfully connected!")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/genres", a.getGenres).Methods("GET")
	a.Router.HandleFunc("/genre", a.createGenre).Methods("POST")
	a.Router.HandleFunc("/genre/{id:[0-9]+}", a.getGenre).Methods("GET")
	a.Router.HandleFunc("/genre/{id:[0-9]+}", a.updateGenre).Methods("PUT")
	a.Router.HandleFunc("/genre/{id:[0-9]+}", a.deleteGenre).Methods("DELETE")
}

func (a *App) getGenres(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	genres, err := genero.GetGenres(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, genres)
}

func (a *App) createGenre(w http.ResponseWriter, r *http.Request) {
	var g genero.Genre
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&g); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := g.CreateGenre(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, g)
}

func (a *App) getGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid genre ID")
		return
	}

	g := genero.Genre{ID: id}
	if err := g.GetGenre(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Genre not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, g)
}

func (a *App) updateGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid genre ID")
		return
	}

	var g genero.Genre
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&g); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	g.ID = id

	if err := g.UpdateGenre(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, g)
}

func (a *App) deleteGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid genre ID")
		return
	}

	g := genero.Genre{ID: id}
	if err := g.DeleteGenre(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
