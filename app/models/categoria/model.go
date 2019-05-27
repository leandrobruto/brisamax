package categoria

import "database/sql"

//t_base_categoria
type Category struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	DataCriacao string `json:"data_criacao"`
}

func (bc *Category) GetBaseCategory(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"nome "+
		"data_criacao FROM t_base_categoria WHERE id=$1",
		bc.ID).Scan(
		&bc.Nome,
		&bc.DataCriacao)
}

func (bc *Category) UpdateBaseCategory(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_base_categoria SET"+
			"nome=$1, "+
			"data_criacao=$2, WHERE id=$3",
			bc.Nome,
			bc.DataCriacao,
			bc.ID)

	return err
}

func (bc *Category) DeleteBaseCategory(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_base_categoria WHERE id=$1", bc.ID)

	return err
}

func (bc *Category) CreateBaseCategory(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_base_categoria("+
			"nome, "+
			"data_criacao) VALUES($1, $2) RETURNING id",
		bc.Nome,
		bc.DataCriacao).Scan(&bc.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetBaseCategories(db *sql.DB, start, count int) ([]Category, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"nome, "+
			"data_criacao FROM t_base_categoria LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	category := []Category{}

	for rows.Next() {
		var bc Category
		if err = rows.Scan(
			&bc.ID,
			&bc.Nome,
			&bc.DataCriacao); err != nil {
			return nil, err
		}
		category = append(category, bc)
	}

	return category, nil
}
