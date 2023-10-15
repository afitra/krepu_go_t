package transaction

import (
	"github.com/labstack/echo"
	"krepu_go_t/models"
)

type Usecase interface {
	UInquiryTransaction(c echo.Context, payload models.PayloadInquiry) (interface{}, error)
	UPayTransactions(c echo.Context, payload models.PayloadPay) (interface{}, error)
}
