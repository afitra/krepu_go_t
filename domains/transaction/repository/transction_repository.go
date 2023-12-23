package repository

import (
	"github.com/jmoiron/sqlx"
	"krepu_go_t/domains/transaction"
	"krepu_go_t/models"
)

type Psql struct {
	sqlx *sqlx.DB
}

func NewPsqlTransaction(sqlx *sqlx.DB) transaction.Repository {
	return &Psql{sqlx}
}

func (p *Psql) RCreateTransaction(payload models.Transaction) error {
	var err error
	query := `
        INSERT INTO transactions (user_id, no_kontrak, otr, admin_fee, cicilan, bunga, nama_asset, tenor, pengajuan)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `

	if _, err = p.sqlx.Exec(query, payload.UserId, payload.NoKontrak, payload.OTR, payload.AdminFee, payload.Cicilan, payload.Bunga, payload.NamaAsset, payload.Tenor, payload.Pengajuan); err != nil {
		return err
	}

	return nil
}

func (p *Psql) RUpdateStatusTransaction(no_kontrak string, status bool) error {
	var err error
	query := " UPDATE transactions SET status = $1 WHERE no_kontrak = $2"
	if _, err = p.sqlx.Exec(query, status, no_kontrak); err != nil {
		return err
	}
	return nil
}
