package GetCurrentTime

import (
	"github.com/labstack/echo"
	"krepu_go_t/models"
)

type Usecase interface {
	UGetCurrentTime(c echo.Context, params models.ParameterTimeZone) (interface{}, error)
}
