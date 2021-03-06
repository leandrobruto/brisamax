package categoria_programa

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"brisamax/app/models/categoria_programa"

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
	a.Router.HandleFunc("/programscategories", a.getCategoryPrograms).Methods("GET")
	a.Router.HandleFunc("/programcategory", a.createCategoryPrograms).Methods("POST")
	a.Router.HandleFunc("/programcategory/{id:[0-9]+}", a.getCategoryProgram).Methods("GET")
	a.Router.HandleFunc("/programcategory/{id:[0-9]+}", a.updateCategoryPrograms).Methods("PUT")
	a.Router.HandleFunc("/programcategory/{id:[0-9]+}", a.deleteCategoryPrograms).Methods("DELETE")
}

/////////////////////////////////////////////////////////
func (a *App) getCategoryPrograms(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	categoryPrograms, err := categoria_programa.GetCategoryPrograms(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, categoryPrograms)
}

func (a *App) createCategoryPrograms(w http.ResponseWriter, r *http.Request) {
	var cp categoria_programa.CategoryProgram
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := cp.CreateCategoryProgram(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, cp)
}

func (a *App) getCategoryProgram(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid program ID")
		return
	}

	cp := categoria_programa.CategoryProgram{ID: id}
	if err := cp.GetCategoryProgram(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Program not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, cp)
}

func (a *App) updateCategoryPrograms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid program ID")
		return
	}

	var cp categoria_programa.CategoryProgram
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	cp.ID = id

	if err := cp.UpdateCategoryProgram(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, cp)
}

func (a *App) deleteCategoryPrograms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid program ID")
		return
	}

	cp := categoria_programa.CategoryProgram{ID: id}
	if err := cp.DeleteCategoryProgram(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

///////////////////////
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
