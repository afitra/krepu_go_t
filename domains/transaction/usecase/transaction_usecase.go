package usecase

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/transaction"
	"krepu_go_t/helpers"
	"krepu_go_t/models"
)

type TransactionUseCase struct {
	response        models.Response
	responseError   models.ErrorResponse
	transactionRepo transaction.Repository
}

func NewTransactionUseCase(repo transaction.Repository) transaction.Usecase {

	return &TransactionUseCase{transactionRepo: repo}
}

func (ut *TransactionUseCase) UInquiryTransaction(c echo.Context, payload models.PayloadInquiry) (interface{}, error) {

	var transaction models.Transaction
	var err error

	for i := 1; i <= 4; i++ {

	}
	user := c.Get("decode").(models.User)
	limit := 0
	switch {
	case payload.Tenor == 1:
		limit = user.TenorSatu
	case payload.Tenor == 2:
		limit = user.TenorDua
	case payload.Tenor == 3:
		limit = user.TenorTiga
	case payload.Tenor == 4:
		limit = user.TenorEmpat

	}

	if payload.Pengajuan > limit {
		return ut.reverseSuccessResponse(models.CodeSuccess, models.ErrorLimitInquiry.Error(), models.ErrSomethingWrong.Error(), nil, models.ErrorLimitInquiry)
	}

	transaction.UserId = user.ID
	transaction.NoKontrak = helpers.GenerateString(10)
	transaction.OTR = payload.Otr
	transaction.AdminFee = payload.AdminFee
	transaction.Cicilan = payload.Cicilan
	transaction.Bunga = payload.Bunga
	transaction.NamaAsset = payload.NamaAsset
	transaction.Tenor = payload.Tenor
	transaction.Pengajuan = payload.Pengajuan

	done := make(chan error)

	go func() {
		if err = ut.transactionRepo.RCreateTransaction(transaction); err != nil {
			done <- err
			return
		}
		done <- nil
	}()

	errFromGoroutine := <-done
	if errFromGoroutine != nil {
		return nil, errFromGoroutine
	}

	return models.ReverseSuccessResponse(models.CodeCreated, models.ResponseSuccess, models.MessageDataProcessing, nil, err)
}

func (ut *TransactionUseCase) UPayTransactions(c echo.Context, payload models.PayloadPay) (interface{}, error) {

	var err error

	if err = ut.transactionRepo.RUpdateStatusTransaction(payload.NoKontrak, true); err != nil {
		return nil, err
	}

	return models.ReverseSuccessResponse(models.CodeSuccess, models.ResponseSuccess, models.MessageDataProcessing, nil, err)
}

func (ut *TransactionUseCase) reverseSuccessResponse(code string, status string, message string, data interface{}, err error) (interface{}, error) {
	var resp models.Response
	if code == "" {
		code = models.CodeSuccess
	}
	resp.Code = code
	resp.Status = status
	resp.Message = message
	resp.Data = data
	return resp, err
}
