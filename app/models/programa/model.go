package programa

import "database/sql"

type Program struct {
	ID                          int    `json:"id"`
	Genero_id                   int    `json:"base_genero_id"`
	Classificacao_indicativa_id int    `json:"base_classificacao_indicativa_id"`
	Titulo                      string `json:"titulo"`
	Url_midia_video             string `json:"url_midia_video"`
	Url_midia_trailer           string `json:"url_midia_trailer"`
	Url_capa_retrato            string `json:"url_capa_retrato"`
	Url_capa_paisagem           string `json:"url_capa_paisagem"`
}

/*
  id bigserial NOT NULL PRIMARY KEY,
  base_genero_id integer NOT NULL,
  base_classificacao_indicativa_id integer NOT NULL,
  titulo character varying(150) NOT NULL,
  url_midia_video character varying(255) NOT NULL,
  url_midia_trailer character varying(255) NOT NULL,
  url_capa_retrato character varying(255) NOT NULL,
  url_capa_paisagem character varying(255) NOT NULL
);
*/

func (p *program) GetProgram(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"gracenote_tms_id, "+
		"gracenote_root_id, "+
		"base_genero_id, "+
		"base_classificacao_indicativa_id, "+
		"titulo, "+
		"ano, "+
		"atores, "+
		"diretores, "+
		"recomendacoes, "+
		"descricao_onga, "+
		"descricao_curta, "+
		"duracao_minutos, "+
		"classificacao_qualidade, "+
		"estrelando, "+
		"destaque, "+
		"episodio_quantidade, "+
		"temporada_quantidade, "+
		"thumbnail_total, "+
		"imagens, "+
		"canal, "+
		"ativo, "+
		"data_disponivel_inicio, "+
		"data_disponivel_fim, "+
		"data_criacao FROM t_programa WHERE id=$1",
		p.ID).Scan(
		&p.GracenoteTmsID,
		&p.GracenoteRootID,
		&p.GeneroID,
		&p.ClassificacaoIndicativaID,
		&p.Titulo,
		&p.Ano,
		&p.Atores,
		&p.Diretores,
		&p.Recomendacoes,
		&p.DescricaoLonga,
		&p.DescricaoCurta,
		&p.DuracaoMinutos,
		&p.ClassificacaoQualidade,
		&p.Estrelando,
		&P.Destaque,
		&p.EpisodioQuantidade,
		&p.TemporadaQuantidade,
		&p.ThumbnailTotal,
		&p.Imagens,
		&p.Canal,
		&p.Ativo,
		&p.DataDisponivelInicio,
		&p.DataDisponivelFim,
		&p.DataCriacao)
}

func (p *program) UpdateProgram(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_programa SET titulo=$1, url_midia_video=$2 WHERE id=$3",
			p.Titulo, p.Url_midia_video, p.ID)

	return err
}

func (p *program) DeleteProgram(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_programa WHERE id=$1", p.ID)

	return err
}

func (p *program) CreateProgram(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_programa("+
			"base_genero_id,"+
			"base_classificacao_indicativa_id,"+
			"titulo,"+
			"url_midia_video,"+
			"url_midia_trailer,"+
			"url_capa_retrato,"+
			"url_capa_paisagem) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		p.Genero_id,
		p.Classificacao_indicativa_id,
		p.Titulo,
		p.Url_midia_video,
		p.Url_midia_trailer,
		p.Url_capa_retrato,
		p.Url_capa_paisagem).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetPrograms(db *sql.DB, start, count int) ([]program, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"base_genero_id, "+
			"base_classificacao_indicativa_id, "+
			"titulo, "+
			"url_midia_video, "+
			"url_midia_trailer, "+
			"url_capa_retrato, "+
			"url_capa_paisagem FROM t_programa LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	programs := []program{}

	for rows.Next() {
		var p program
		if err = rows.Scan(
			&p.ID,
			&p.Genero_id,
			&p.Classificacao_indicativa_id,
			&p.Titulo,
			&p.Url_midia_video,
			&p.Url_midia_trailer,
			&p.Url_capa_retrato,
			&p.Url_capa_paisagem); err != nil {
			return nil, err
		}
		programs = append(programs, p)
	}

	return programs, nil
}
