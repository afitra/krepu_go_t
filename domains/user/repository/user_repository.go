package repository

import (
	"github.com/jmoiron/sqlx"
	"krepu_go_t/domains/user"
)

type PsqlUser struct {
	sqlx *sqlx.DB
}

func NewPsqlUniqLink(sqlx *sqlx.DB) user.Repository {
	return &PsqlUser{sqlx}
}
