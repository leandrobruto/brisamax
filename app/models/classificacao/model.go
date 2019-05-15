package classificacao

import "database/sql"

type Classification struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Nivel       int    `json:"nivel"`
	DataCriacao string `json:"data_criacao`
}

func (c *Classification) GetClassification(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"nome, "+
		"nivel, "+
		"data_criacao FROM t_base_classificacao_indicativa WHERE if=$1",
		c.ID).Scan(
		&c.Nome,
		&c.Nivel,
		&c.DataCriacao)
}

func (c *Classification) UpdateClassification(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_base_classificacao_indicativa SET"+
			"nome=$1, "+
			"nivel=$2, "+
			"data_criacao=$3, WHERE id=$4",
			c.Nome,
			c.Nivel,
			c.DataCriacao,
			c.ID)

	return err
}

func (c *Classification) DeleteClassification(db *sql.DB) error {
	_, err :=
		db.Exec("DELETE FROM t_base_classificacao_indicativa WHERE id=$1",
			c.ID)

	return err
}

func (c *Classification) CreateClassification(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_base_classificacao_indicativa ("+
			"nome, "+
			"nivel, "+
			"data_criacao) VALUES ($1,$2,$3) RETURNING id",
		c.Nome,
		c.Nivel,
		c.DataCriacao).Scan(&c.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetRatings(db *sql.DB, start, count int) ([]Classification, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"nome, "+
			"nivel, "+
			"data_criacao FROM t_base_classificacao_indicativa LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ratings := []Classification{}

	for rows.Next() {
		var c Classification
		if err = rows.Scan(
			&c.ID,
			&c.Nome,
			&c.Nivel,
			&c.DataCriacao); err != nil {
			return nil, err
		}
		ratings = append(ratings, c)
	}

	return ratings, nil
}
