package serie

import "database/sql"

// id bigint NOT NULL DEFAULT nextval('t_programa_serie_id_seq'::regclass),
//   programa_id bigint NOT NULL,
//   nome_serie character varying(150) NOT NULL,
//   temporada_numero integer NOT NULL,
//   serie_numero integer NOT NULL,
//   descricao character varying(250),
//   imagens jsonb,
//   url_midia_video character varying(255) NOT NULL,
//   url_midia_trailer character varying(255) NOT NULL,
//   url_capa_retrato character varying(255) NOT NULL,
//   url_capa_paisagem character varying(255) NOT NULL,
//   data_criacao timestamp(0) without time zone NOT NULL DEFAULT now(),

type Series struct {
	ID              int    `json:"id"`
	ProgramaID      string `json:"programa_id"`
	NomeSerie       string `json:"nome_serie"`
	TemporadaNumero int    `json:"temporada_numero"`
	SerieNumero     int    `json:"serie_numero"`
	Descricao       string `json:"descricao"`
	Imagens         string `json:"imagens"` //jsonb
	UrlMidiaVideo   string `json:"url_midia_video"`
	UrlMidiaTrailer string `json:"url_midia_trailer"`
	UrlCapaRetrato  string `json:"url_capa_retrato"`
	UrlCapaPaisagem string `json:"url_capa_paisagem"`
	DataCriacao     string `json:"data_criacao"`
}

func (s *Series) GetSerie(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"programa_id, "+
		"nome_serie, "+
		"temporada_numero, "+
		"serie_numero, "+
		"descricao, "+
		"imagens, "+
		"url_midia_video, "+
		"url_midia_trailer, "+
		"url_capa_retrato, "+
		"url_capa_paisagem, "+
		"data_criacao FROM t_programa_serie	WHERE id=$1",
		s.ID).Scan(
		&s.ProgramaID,
		&s.NomeSerie,
		&s.TemporadaNumero,
		&s.SerieNumero,
		&s.Descricao,
		&s.Imagens,
		&s.UrlMidiaVideo,
		&s.UrlMidiaTrailer,
		&s.UrlCapaRetrato,
		&s.UrlCapaPaisagem,
		&s.DataCriacao)
}

func (s *Series) UpdateSeries(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_programa_serie SET"+
			"programa_id=$1, "+
			"nome_serie=$2, "+
			"temporada_numero=$3, "+
			"serie_numero=$4, "+
			"descricao=$5, "+
			"imagens=$6, "+
			"url_midia_video=$7, "+
			"url_midia_trailer=$8, "+
			"url_capa_retrato=$9, "+
			"url_capa_paisagem=$10, "+
			"data_criacao=$11,, WHERE id=$12",
			s.ProgramaID,
			s.NomeSerie,
			s.TemporadaNumero,
			s.SerieNumero,
			s.Descricao,
			s.Imagens,
			s.UrlMidiaVideo,
			s.UrlMidiaTrailer,
			s.UrlCapaRetrato,
			s.UrlCapaPaisagem,
			s.DataCriacao,
			s.ID)

	return err
}

func (s *Series) DeleteSeries(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_programa_serie WHERE id=$1", s.ID)

	return err
}

func (s *Series) CreateSeries(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_programa_serie("+
			"programa_id, "+
			"nome_serie, "+
			"temporada_numero, "+
			"serie_numero, "+
			"descricao, "+
			"imagens, "+
			"url_midia_video, "+
			"url_midia_trailer, "+
			"url_capa_retrato, "+
			"url_capa_paisagem, "+
			"data_criacao) VALUES($1, $2, $3, $4, $5, $6, $7, &8,"+
			"&9, &10, &11) RETURNING id",
		s.ProgramaID,
		s.NomeSerie,
		s.TemporadaNumero,
		s.SerieNumero,
		s.Descricao,
		s.Imagens,
		s.UrlMidiaVideo,
		s.UrlMidiaTrailer,
		s.UrlCapaRetrato,
		s.UrlCapaPaisagem,
		s.DataCriacao).Scan(&s.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetSeries(db *sql.DB, start, count int) ([]Series, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"programa_id, "+
			"nome_serie, "+
			"temporada_numero, "+
			"serie_numero, "+
			"descricao, "+
			"imagens, "+
			"url_midia_video, "+
			"url_midia_trailer, "+
			"url_capa_retrato, "+
			"url_capa_paisagem, "+
			"data_criacao FROM t_programa_serie LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	series := []Series{}

	for rows.Next() {
		var s Series
		if err = rows.Scan(
			&s.ID,
			&s.ProgramaID,
			&s.NomeSerie,
			&s.TemporadaNumero,
			&s.SerieNumero,
			&s.Descricao,
			&s.Imagens,
			&s.UrlMidiaVideo,
			&s.UrlMidiaTrailer,
			&s.UrlCapaRetrato,
			&s.UrlCapaPaisagem,
			&s.DataCriacao); err != nil {
			return nil, err
		}
		series = append(series, s)
	}

	return series, nil
}
