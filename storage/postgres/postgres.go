package postgres

import (
	"e-wallet/storage/models"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}


func (d Database) GetBalance(w models.Wallet) (*models.Wallet,error) {

	query := `select balance where wallet_id = $1 and delted_at is null`

	err := d.db.QueryRow(query,w.Id).Scan(&w.Balance)
	if err != nil {
		return nil,err
	}
	return &models.Wallet{w.Id,w.Balance},nil
}
