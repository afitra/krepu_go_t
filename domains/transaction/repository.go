package transaction

import (
	"krepu_go_t/models"
)

type Repository interface {
	RCreateTransaction(payload models.Transaction) error
	RUpdateStatusTransaction(no_kontrak string, status bool) error
}
