package http

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/transaction"
	"krepu_go_t/logger"
	"krepu_go_t/models"
	"net/http"
)

type TransactionHandler struct {
	response   models.Response
	respErrors models.ErrorResponse
	usecase    transaction.Usecase
}

func NewTransactionHandler(echoGroup models.EchoGroup, tuc transaction.Usecase) {
	handler := &TransactionHandler{
		usecase: tuc,
	}
	echoGroup.AUTH.POST("/transaction/inquiry", handler.inquiryTransaction)
	echoGroup.ADMIN.POST("/transaction/pay", handler.payTransaction)
}

func (th *TransactionHandler) inquiryTransaction(c echo.Context) error {
	var request models.PayloadInquiry
	th.response, th.respErrors = models.NewResponse()
	if err := c.Bind(&request); err != nil {

		logger.Make(c, nil).Debug(err)
		th.respErrors.SetTitle(models.MessageUnprocessableEntity)
		th.response.SetResponse("", &th.respErrors)
		return th.response.Body(c, err)
	}

	if err := c.Validate(request); err != nil {
		logger.Make(c, nil).Debug(err)
		th.respErrors.SetTitle(models.ErrSomethingWrong.Error())
		th.respErrors.AddError(err.Error())
		th.response.SetResponse("", &th.respErrors)
		return th.response.Body(c, err)
	}

	resp, err := th.usecase.UInquiryTransaction(c, request)
	if err != nil {
		logger.Make(c, nil).Debug(err)
		th.respErrors.SetTitle(err.Error())
		th.response.SetResponse("", &th.respErrors)
		return th.response.Body(c, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (th *TransactionHandler) payTransaction(c echo.Context) error {
	var request models.PayloadPay
	th.response, th.respErrors = models.NewResponse()
	if err := c.Bind(&request); err != nil {

		logger.Make(c, nil).Debug(err)
		th.respErrors.SetTitle(models.MessageUnprocessableEntity)
		th.response.SetResponse("", &th.respErrors)
		return th.response.Body(c, err)
	}

	if err := c.Validate(request); err != nil {
		logger.Make(c, nil).Debug(err)
		th.respErrors.SetTitle(models.ErrSomethingWrong.Error())
		th.respErrors.AddError(err.Error())
		th.response.SetResponse("", &th.respErrors)
		return th.response.Body(c, err)
	}

	resp, err := th.usecase.UPayTransactions(c, request)
	if err != nil {
		logger.Make(c, nil).Debug(err)
		th.respErrors.SetTitle(err.Error())
		th.response.SetResponse("", &th.respErrors)
		return th.response.Body(c, err)
	}
	return c.JSON(http.StatusOK, resp)
}
