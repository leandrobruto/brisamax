package genero

import "database/sql"

type Genre struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	DataCriacao string `json:"data_criacao"`
}

func (g *Genre) GetGenre(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"nome, "+
		"data_criacao FROM t_base_genero WHERE id=$1", //ver esse id=$1 com leandro
		g.ID).Scan(
		&g.Nome,
		&g.DataCriacao)
}

func (g *Genre) UpdateGenre(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_base_genero SET"+
			"nome=$1, "+
			"data_criacao=$2, WHERE id=$3",
			g.Nome,
			g.DataCriacao,
			g.ID)

	return err
}

func (g *Genre) DeleteGenre(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_programa WHERE id=$1", g.ID)

	return err
}

func (g *Genre) CreateGenre(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_base_genero("+
			"nome, "+
			"data_criacao) VALUES($1,$2)RETURNING id",
		g.Nome,
		g.DataCriacao).Scan(&g.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetGenres(db *sql.DB, start, count int) ([]Genre, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"nome, "+
			"data_criacao FROM t_base_genero LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := []Genre{}

	for rows.Next() {
		var g Genre
		if err = rows.Scan(
			&g.ID,
			&g.Nome,
			&g.DataCriacao); err != nil {
			return nil, err
		}
		genres = append(genres, g)
	}

	return genres, nil
}
