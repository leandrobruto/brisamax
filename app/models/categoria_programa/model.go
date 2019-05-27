package categoria_programa

import "database/sql"

type CategoryProgram struct {
	ID              int    `json:"id"`
	ProgramaID      string `json:"programa_id"`
	BaseCategoriaID int    `json:"base_categoria_id"`
	DataCriacao     string `json:"data_criacao"`
}

func (cp *CategoryProgram) GetCategoryProgram(db *sql.DB) error {
	return db.QueryRow("SELECT "+
		"programa_id, "+
		"base_categoria_id, "+
		"data_criacao FROM t_programa_categoria WHERE id=$1",
		cp.ID).Scan(
		&cp.ProgramaID,
		&cp.BaseCategoriaID,
		&cp.DataCriacao)
}
func (cp *CategoryProgram) UpdateCategoryProgram(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE t_programa_categoria SET"+
			"programa_id=$1, "+
			"base_categoria_id=$2, "+
			"data_criacao=$3, WHERE id=$4",
			cp.ProgramaID,
			cp.BaseCategoriaID,
			cp.DataCriacao,
			cp.ID)

	return err
}

func (cp *CategoryProgram) DeleteCategoryProgram(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM t_programa_categoria WHERE id=$1", cp.ID)

	return err
}

func (cp *CategoryProgram) CreateCategoryProgram(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO t_programa_categoria("+
			"programa_id, "+
			"base_categoria_id, "+
			"data_criacao) VALUES($1, $2, $3) RETURNING id",
		cp.ProgramaID,
		cp.BaseCategoriaID,
		cp.DataCriacao).Scan(&cp.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetCategoryPrograms(db *sql.DB, start, count int) ([]CategoryProgram, error) {
	rows, err := db.Query(
		"SELECT "+
			"id, "+
			"programa_id, "+
			"base_categoria_id, "+
			"data_criacao FROM t_programa_categoria LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categoryPrograms := []CategoryProgram{}

	for rows.Next() {
		var cp CategoryProgram
		if err = rows.Scan(
			&cp.ID,
			&cp.ProgramaID,
			&cp.BaseCategoriaID,
			&cp.DataCriacao); err != nil {
			return nil, err
		}
		categoryPrograms = append(categoryPrograms, cp)
	}

	return categoryPrograms, nil
}
