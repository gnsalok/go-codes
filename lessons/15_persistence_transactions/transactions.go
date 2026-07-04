package persistencetransactions

import (
	"context"
	"database/sql"
)

type Beginner struct {
	ID   int64
	Name string
}

func CreateBeginner(ctx context.Context, db *sql.DB, name string) (int64, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, `INSERT INTO beginners(name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}
