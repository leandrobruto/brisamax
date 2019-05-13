package programa

import "database/sql"

type Program struct {
	ID                            int    `json:"id"`
	GracenoteTmsID                string `json:gracenote_tms_id`
	GracenoteRootID               string `json:gracenote_root_id`
	BaseGeneroID                  int    `json:"base_genero_id"`
	BaseClassificacaoIndicativaID int    `json:"base_classificacao_indicativa_id"`
	Titulo                        string `json:"titulo"`
	Ano                           int    `json:ano`
	Atores                        string `json:atores`        //jsonb
	Diretores                     string `json:diretores`     //jsonb
	Recomendacoes                 string `json:recomendacoes` //jsonb
	DescricaoLonga                string `json:descricao_longa`
	DescricaoCurta                string `json:descricao_curta`
	DuracaoMinutos                int    `json:duracao_minutos`
	ClassificacaoQualidade        string `json:classificacao_qualidade`
	Estrelando                    string `json:estrelando` //jsonb
	Destaque                      bool   `json:destaque`
	EpisodioQuantidade            int    `json:episodeo_quantidade`
	TemporadaQuantidade           int    `json:temporada_quantidade`
	ThumbnailTotal                int    `json:thumbnail_total`
	Imagens                       string `json:imagens` //jsonb
	Canal                         bool   `json:canal`
	Ativo                         bool   `json:ativo`
	DataDisponivelInicio          string `json:data_disponivel_inicio`
	DataDisponivelFim             string `json:data_disponivel_fim`
	DataCriacao                   string `json:data_criacao`
}

/*
  id bigserial NOT NULL PRIMARY KEY,
  gracenote_tms_id character varying(100),
  gracenote_root_id bigint,
  base_genero_id integer NOT NULL,
  base_classificacao_indicativa_id integer NOT NULL,
  titulo character varying(150) NOT NULL,
  ano integer,
  atores jsonb,
  diretores jsonb,
  recomendacoes jsonb,
  descricao_longa text,
  descricao_curta text,
  duracao_minutos integer,
  classificacao_qualidade numeric,
  estrelando jsonb,
  destaque boolean NOT NULL DEFAULT false,
  episodio_quantidade integer,
  temporada_quantidade integer,
  thumbnail_total integer,
  imagens jsonb,
  canal boolean NOT NULL DEFAULT false,
  ativo boolean NOT NULL DEFAULT false,
  data_disponivel_inicio timestamp(0) without time zone NOT NULL DEFAULT NOW(),
  data_disponivel_fim timestamp(0) without time zone NOT NULL DEFAULT NOW(),
  data_criacao timestamp(0) without time zone NOT NULL DEFAULT NOW()
);
*/

func (p *Program) GetProgram(db *sql.DB) error {
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
		"descricao_longa, "+
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
		&p.BaseGeneroID,
		&p.BaseClassificacaoIndicativaID,
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
		&p.Destaque,
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

func (p *Program) UpdateProgram(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_programa SET"+
			"gracenote_tms_id=$1, "+
			"gracenote_root_id=$2, "+
			"base_genero_id=$3, "+
			"base_classificacao_indicativa_id=$4, "+
			"titulo=$5, "+
			"ano=$6, "+
			"atores=$7, "+
			"diretores=$8, "+
			"recomendacoes=$9, "+
			"descricao_longa=$10, "+
			"descricao_curta=$11, "+
			"duracao_minutos=$12, "+
			"classificacao_qualidade=$13, "+
			"estrelando=$14, "+
			"destaque=$15, "+
			"episodio_quantidade=$16, "+
			"temporada_quantidade, "+
			"thumbnail_total=$17, "+
			"imagens=$18, "+
			"canal=$19, "+
			"ativo=$20, "+
			"data_disponivel_inicio=$21, "+
			"data_disponivel_fim=$22, WHERE id=$23",
			p.GracenoteTmsID,
			p.GracenoteRootID,
			p.BaseGeneroID,
			p.BaseClassificacaoIndicativaID,
			p.Titulo,
			p.Ano,
			p.Atores,
			p.Diretores,
			p.Recomendacoes,
			p.DescricaoLonga,
			p.DescricaoCurta,
			p.DuracaoMinutos,
			p.ClassificacaoQualidade,
			p.Estrelando,
			p.Destaque,
			p.EpisodioQuantidade,
			p.TemporadaQuantidade,
			p.ThumbnailTotal,
			p.Imagens,
			p.Canal,
			p.Ativo,
			p.DataDisponivelInicio,
			p.DataDisponivelFim,
			p.ID)

	return err
}

func (p *Program) DeleteProgram(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_programa WHERE id=$1", p.ID)

	return err
}

func (p *Program) CreateProgram(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_programa("+
			"gracenote_tms_id, "+
			"gracenote_root_id, "+
			"base_genero_id, "+
			"base_classificacao_indicativa_id, "+
			"titulo, "+
			"ano, "+
			"atores, "+
			"diretores, "+
			"recomendacoes, "+
			"descricao_longa, "+
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
			"data_disponivel_fim) VALUES($1, $2, $3, $4, $5, $6, $7, &8,"+
			"&9, &10, &11, &12, &13, &14, &15, &16, &17, &18, &19, &20,"+
			"&21, &22, &23) RETURNING id",
		p.GracenoteTmsID,
		p.GracenoteRootID,
		p.BaseGeneroID,
		p.BaseClassificacaoIndicativaID,
		p.Titulo,
		p.Ano,
		p.Atores,
		p.Diretores,
		p.Recomendacoes,
		p.DescricaoLonga,
		p.DescricaoCurta,
		p.DuracaoMinutos,
		p.ClassificacaoQualidade,
		p.Estrelando,
		p.Destaque,
		p.EpisodioQuantidade,
		p.TemporadaQuantidade,
		p.ThumbnailTotal,
		p.Imagens,
		p.Canal,
		p.Ativo,
		p.DataDisponivelInicio,
		p.DataDisponivelFim,
		p.DataCriacao).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetPrograms(db *sql.DB, start, count int) ([]Program, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"gracenote_tms_id, "+
			"gracenote_root_id, "+
			"base_genero_id, "+
			"base_classificacao_indicativa_id, "+
			"titulo, "+
			"ano, "+
			"atores, "+
			"diretores, "+
			"recomendacoes, "+
			"descricao_longa, "+
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
			"data_criacao FROM t_programa LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	programs := []Program{}

	for rows.Next() {
		var p Program
		if err = rows.Scan(
			&p.ID,
			&p.GracenoteTmsID,
			&p.GracenoteRootID,
			&p.BaseGeneroID,
			&p.BaseClassificacaoIndicativaID,
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
			&p.Destaque,
			&p.EpisodioQuantidade,
			&p.TemporadaQuantidade,
			&p.ThumbnailTotal,
			&p.Imagens,
			&p.Canal,
			&p.Ativo,
			&p.DataDisponivelInicio,
			&p.DataDisponivelFim,
			&p.DataCriacao); err != nil {
			return nil, err
		}
		programs = append(programs, p)
	}

	return programs, nil
}
