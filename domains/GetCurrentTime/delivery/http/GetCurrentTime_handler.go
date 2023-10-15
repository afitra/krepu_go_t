package http

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/GetCurrentTime"
	"krepu_go_t/logger"
	"krepu_go_t/models"
	"net/http"
)

type GetCurrentTimeHandler struct {
	response   models.Response
	respErrors models.ErrorResponse
	usecase    GetCurrentTime.Usecase
}

func NewGetCurrentTimeHandler(echoGroup models.EchoGroup, auc GetCurrentTime.Usecase) {
	handler := &GetCurrentTimeHandler{
		usecase: auc,
	}
	echoGroup.API.GET("/GetCurrentTime", handler.getTimeData)
}

func (gth *GetCurrentTimeHandler) getTimeData(c echo.Context) error {
	var err error
	var resp interface{}
	var request models.ParameterTimeZone
	request.TimeZone = c.QueryParam("timeZone")

	if err = c.Validate(request); err != nil {
		errorResponse := setDefaultError()
		errorResponse.Errors["timeZone"] = append(errorResponse.Errors["timeZone"], "The timeZone field is required.")
		logger.Make(c, nil).Debug(err)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if resp, err = gth.usecase.UGetCurrentTime(c, request); err != nil {
		logger.Make(c, nil).Debug(err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid timeZone")
	}

	return c.JSON(http.StatusOK, resp)
}

func setDefaultError() models.ErrorResponse {
	var errorResponse models.ErrorResponse
	errorResponse.Type = "https://tools.ietf.org/html/rfc7231#section-6.5.1"
	errorResponse.Title = "One or more validation errors occurred."
	errorResponse.TraceID = "00-ac218d4c7be84b1490cf84c52aa546ba-938ba58092724fb4-00"
	errorResponse.Status = http.StatusBadRequest
	errorResponse.Errors = map[string][]string{}
	return errorResponse
}
