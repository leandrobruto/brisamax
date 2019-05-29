package serie_programa

import "database/sql"

//t_programa_serie
type SeriePrograma struct {
	ID              int    `json:"id"`
	ProgramaID      int    `json:"programa_id"`
	NomeSerie       string `json:"nome_serie"`
	TemporadaNumero int    `json:"temporada_numero"`
	SerieNumero     int    `json:"serie_numero"`
	Descricao       string `json:"descricao"`
	Imagens         string `json:"imagens"`
	DataCriacao     string `json:"data_criacao"`
}

func (sp *SeriePrograma) GetSeriePrograma(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"programa_id "+
		"nome_serie "+
		"temporada_numero "+
		"serie_numero "+
		"descricao "+
		"imagens"+
		"data_criacao FROM t_programa_serie WHERE id=$1",
		sp.ID).Scan(
		&sp.ProgramaID,
		&sp.NomeSerie,
		&sp.TemporadaNumero,
		&sp.SerieNumero,
		&sp.Descricao,
		&sp.Imagens,
		&sp.DataCriacao)
}

func (sp *SeriePrograma) UpdateSeriePrograma(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_programa_serie SET"+
			"programa_id=$1"+
			"nome_serie=$2"+
			"temporada_numero=$3"+
			"serie_numero=$4"+
			"descricao=$5"+
			"imagens=$6"+
			"data_criacao=$7, WHERE id=$8",
			sp.ProgramaID,
			sp.NomeSerie,
			sp.TemporadaNumero,
			sp.SerieNumero,
			sp.Descricao,
			sp.Imagens,
			sp.DataCriacao,
			sp.ID)

	return err
}

func (sp *SeriePrograma) DeleteSeriePrograma(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_programa_serie WHERE id=$1", sp.ID)

	return err
}

func (sp *SeriePrograma) CreateSeriePrograma(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_programa_serie("+
			"programa_id,"+
			"nome_serie,"+
			"temporada_numero,"+
			"serie_numero,"+
			"descricao,"+
			"imagens,"+
			"data_criacao) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		sp.ProgramaID,
		sp.NomeSerie,
		sp.TemporadaNumero,
		sp.SerieNumero,
		sp.Descricao,
		sp.Imagens,
		sp.DataCriacao).Scan(&sp.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetSeriePrograms(db *sql.DB, start, count int) ([]SeriePrograma, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"programa_id "+
			"nome_serie "+
			"temporada_numero "+
			"serie_numero "+
			"descricao "+
			"imagens"+
			"data_criacao FROM t_programa_serie LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	seriePrograma := []SeriePrograma{}

	for rows.Next() {
		var sp SeriePrograma
		if err = rows.Scan(
			&sp.ID,
			&sp.ProgramaID,
			&sp.NomeSerie,
			&sp.TemporadaNumero,
			&sp.SerieNumero,
			&sp.Descricao,
			&sp.Imagens,
			&sp.DataCriacao); err != nil {
			return nil, err
		}
		seriePrograma = append(seriePrograma, sp)
	}

	return seriePrograma, nil
}
