package models

import (
	"github.com/labstack/echo"
	"net/http"
	"reflect"
	"strings"
)

var (
	ResponseSuccess = "success"
	ResponseError   = "error"
	CodeSuccess     = "200"
	CodeCreated     = "201"
	CodeBad         = "400"
)

var responseCode = map[string]string{
	ResponseSuccess: "00",
	ResponseError:   "99",
}

func (re *ErrorResponse) SetTitle(title string) {
	re.Title = title
}

func (re *ErrorResponse) SetTitleCode(code string, title string, desc string) {
	re.Title = title
	re.Code = code
	re.AddError(desc)
}

type Response struct {
	Code        string      `json:"code,omitempty"`
	Status      string      `json:"status,omitempty"`
	Message     string      `json:"message,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Title   string
	Code    string
	Message string
	Details []string
}

type ResponseErrorData struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewResponse() (Response, ErrorResponse) {
	return Response{}, ErrorResponse{}
}

func (re *ErrorResponse) AddError(errString string) {
	re.Details = append(re.Details, errString)
}

func (resp *Response) SetResponse(respData interface{}, respErrors *ErrorResponse) {
	typeResp := reflect.TypeOf(respData)

	if typeResp.Kind() != reflect.Slice {
		resp.Data = respData
	}

	if respErrors.Title == "" {
		resp.Status = ResponseSuccess
		resp.Code = responseCode[ResponseSuccess]
		resp.Message = MessageDataSuccess
		resp.Data = respData

		return
	}

	resp.Status = ResponseError
	resp.Code = responseCode[ResponseError]
	resp.Message = respErrors.Title

	if len(respErrors.Details) != 0 {
		resp.Description = strings.Join(respErrors.Details, ", ")
	}

	if respErrors.Code != "" {
		resp.Code = respErrors.Code
	}
}

func (resp *Response) Body(c echo.Context, err error) error {
	return c.JSON(getStatusCode(err), resp)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if strings.Contains(err.Error(), "400") {
		return http.StatusBadRequest
	}

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrUnauthorized, ErrPassword:
		return http.StatusUnauthorized
	default:
		return http.StatusOK
	}
}
