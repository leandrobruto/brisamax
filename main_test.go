package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"brisamax/app/controllers/programa"
)

var a programa.App

func TestMain(m *testing.M) {
	a = programa.App{}
	a.Initialize(
		os.Getenv("postgres"),
		os.Getenv("admin"),
		os.Getenv("brisamax"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/programs", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentProgram(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/program/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Program not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Program not found'. Got '%s'", m["error"])
	}
}

func TestCreateProgram(t *testing.T) {
	clearTable()
	addGenero(1)
	addClassificacaoIndicativa(1)

	payload := []byte(`{"base_genero_id": 1, "base_classificacao_indicativa_id": 1, "titulo":"titulo", "url_midia_video":"url_midia_video", "url_midia_trailer":"url_midia_trailer", "url_capa_retrato":"url_capa_retrato", "url_capa_paisagem":"url_capa_paisagem"}`)

	req, _ := http.NewRequest("POST", "/program", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["base_genero_id"] != 1.0 {
		t.Errorf("Expected program base_genero_id to be '1'. Got '%v'", m["base_genero_id"])
	}

	if m["base_classificacao_indicativa_id"] != 1.0 {
		t.Errorf("Expected program base_classificacao_indicativa_id to be '1'. Got '%v'", m["base_classificacao_indicativa_id"])
	}

	if m["titulo"] != "titulo" {
		t.Errorf("Expected program title to be 'titulo'. Got '%v'", m["titulo"])
	}

	if m["url_midia_video"] != "url_midia_video" {
		t.Errorf("Expected product url_midia_video to be 'url_midia_video'. Got '%v'", m["url_midia_video"])
	}

	if m["url_midia_trailer"] != "url_midia_trailer" {
		t.Errorf("Expected program url_midia_trailer to be 'url_midia_trailer'. Got '%v'", m["url_midia_trailer"])
	}

	if m["url_capa_retrato"] != "url_capa_retrato" {
		t.Errorf("Expected program url_capa_retrato to be 'url_capa_retrato'. Got '%v'", m["url_capa_retrato"])
	}

	if m["url_capa_paisagem"] != "url_capa_paisagem" {
		t.Errorf("Expected program url_capa_paisagem to be 'url_capa_paisagem'. Got '%v'", m["url_capa_paisagem"])
	}
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected program ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetProgram(t *testing.T) {
	clearTable()
	addGenero(1)
	addClassificacaoIndicativa(1)
	addPrograms(1)

	req, _ := http.NewRequest("GET", "/program/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateProgram(t *testing.T) {
	clearTable()
	addGenero(1)
	addClassificacaoIndicativa(1)
	addPrograms(1)

	req, _ := http.NewRequest("GET", "/program/1", nil)
	response := executeRequest(req)
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	payload := []byte(`{"titulo":"titulo_updated", "url_midia_video":"url_midia_video_updated"}`)

	req, _ = http.NewRequest("PUT", "/program/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalProduct["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["id"], m["id"])
	}

	if m["titulo"] == originalProduct["titulo"] {
		t.Errorf("Expected the titulo to change from '%v' to '%v'. Got '%v'", originalProduct["titulo"], m["titulo"], m["titulo"])
	}

	if m["url_midia_video"] == originalProduct["url_midia_video"] {
		t.Errorf("Expected the url_midia_video to change from '%v' to '%v'. Got '%v'", originalProduct["url_midia_video"], m["url_midia_video"], m["url_midia_video"])
	}
}

func TestDeleteProgram(t *testing.T) {
	clearTable()
	addGenero(1)
	addClassificacaoIndicativa(1)
	addPrograms(1)

	req, _ := http.NewRequest("GET", "/program/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/program/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/program/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM t_programa")
	a.DB.Exec("ALTER SEQUENCE t_programa_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM t_base_genero")
	a.DB.Exec("ALTER SEQUENCE t_base_genero_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM t_base_classificacao_indicativa")
	a.DB.Exec("ALTER SEQUENCE t_base_classificacao_indicativa_id_seq RESTART WITH 1")
}

func addGenero(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO t_base_genero (nome) VALUES ($1)", "Genero "+strconv.Itoa(i))
	}
}

func addClassificacaoIndicativa(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO t_base_classificacao_indicativa(nome, nivel) VALUES($1, $2)", "Classificacao "+strconv.Itoa(i), (i+1.0)*10)
	}
}

func addPrograms(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO t_programa(base_genero_id, base_classificacao_indicativa_id, titulo, url_midia_video, url_midia_trailer, url_capa_retrato, url_capa_paisagem) VALUES($1, $2, $3, $4, $5, $6, $7)", 1, 1, "Product "+strconv.Itoa(i), "url_midia_video", "url_midia_trailer", "url_capa_retrato", "url_capa_paisagem")
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS t_programa
(
    id bigserial NOT NULL PRIMARY KEY,
	base_genero_id integer NOT NULL,
	base_classificacao_indicativa_id integer NOT NULL,
	titulo character varying(150) NOT NULL,
	url_midia_video character varying(255) NOT NULL,
	url_midia_trailer character varying(255) NOT NULL,
	url_capa_retrato character varying(255) NOT NULL,
	url_capa_paisagem character varying(255) NOT NULL
)`