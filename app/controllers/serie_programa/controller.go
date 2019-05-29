package serie_programa

import (
	"brisamax/app/models/serie_programa"
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
		//"user=postgres password=admin dbname=brisamax sslmode=disable"

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	fmt.Println("Successfully connected!")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/seriesProgramas", a.getSeriesPrograms).Methods("GET")
	a.Router.HandleFunc("/seriePrograma", a.createSerieProgram).Methods("POST")
	a.Router.HandleFunc("/seriePrograma/{id:[0-9]+}", a.getSerieProgram).Methods("GET")
	a.Router.HandleFunc("/seriePrograma/{id:[0-9]+}", a.updateSerieProgram).Methods("PUT")
	a.Router.HandleFunc("/seriePrograma/{id:[0-9]+}", a.deleteSerieProgram).Methods("DELETE")
}

func (a *App) getSeriesPrograms(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	seriePrograma, err := serie_programa.GetSeriePrograms(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, seriePrograma)
}

func (a *App) createSerieProgram(w http.ResponseWriter, r *http.Request) {
	var sp serie_programa.SeriePrograma
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := sp.CreateSeriePrograma(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, sp)
}

func (a *App) getSerieProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid program ID")
		return
	}

	sp := serie_programa.SeriePrograma{ID: id}
	if err := sp.GetSeriePrograma(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Program not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, sp)
}

func (a *App) updateSerieProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid program ID")
		return
	}

	var sp serie_programa.SeriePrograma
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	sp.ID = id

	if err := sp.UpdateSeriePrograma(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, sp)
}

func (a *App) deleteSerieProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid program ID")
		return
	}

	sp := serie_programa.SeriePrograma{ID: id}
	if err := sp.DeleteSeriePrograma(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//////////////////////////
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
