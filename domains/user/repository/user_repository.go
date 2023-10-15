package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"krepu_go_t/domains/user"
	"krepu_go_t/models"
)

type Psql struct {
	sqlx *sqlx.DB
}

func NewPsqlUniqLink(sqlx *sqlx.DB) user.Repository {
	return &Psql{sqlx}
}

func (p *Psql) RCreateUser(c echo.Context, payload models.UserPayload) error {
	query := "INSERT INTO users (nik, user_name, password, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	_, err := p.sqlx.Exec(query, payload.Nik, payload.UserName, payload.Password, payload.FullName, payload.LegalName, payload.TempatLahir, payload.TanggalLahir, payload.Gaji, payload.FotoKTP, payload.FotoSelfie)
	if err != nil {
		return err
	}

	return nil
}

func (p *Psql) RGetUserByUserName(c echo.Context, user_name string) (models.User, error) {

	var user models.User
	query := "SELECT * FROM users WHERE user_name = $1"
	err := p.sqlx.Get(&user, query, user_name)
	if err != nil {
		return user, err
	}
	return user, nil
}