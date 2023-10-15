package usecase

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/GetCurrentTime"
	"krepu_go_t/models"
	"time"
)

type GetCurrentTimeUseCase struct {
	response models.Response
}

func NewGetCurrentTimeUseCase() GetCurrentTime.Usecase {

	return &GetCurrentTimeUseCase{}
}

func (gcr *GetCurrentTimeUseCase) UGetCurrentTime(c echo.Context, params models.ParameterTimeZone) (interface{}, error) {
	var resp models.Response
	location, err := time.LoadLocation(params.TimeZone)
	if err != nil {
		return resp, err
	}

	currentTime := time.Now().In(location)

	gcr.response.Year = currentTime.Year()
	gcr.response.Month = int(currentTime.Month())
	gcr.response.Day = currentTime.Day()
	gcr.response.Hour = currentTime.Hour()
	gcr.response.Minute = currentTime.Minute()
	gcr.response.Seconds = currentTime.Second()
	gcr.response.MilliSeconds = currentTime.Nanosecond() / 1000000
	gcr.response.DateTime = currentTime.Format("2006-01-02T15:04:05.9999999")
	gcr.response.Date = currentTime.Format("01/02/2006")
	gcr.response.Time = currentTime.Format("15:04")
	gcr.response.TimeZone = currentTime.Location().String()
	gcr.response.DayOfWeek = currentTime.Weekday().String()
	gcr.response.DstActive = location != time.FixedZone(params.TimeZone, 0)

	return gcr.response, nil
}
