package models

// Response struct is represent a data for output payload
type Response struct {
	Year         int    `json:"year,omitempty"`
	Month        int    `json:"month,omitempty"`
	Day          int    `json:"day,omitempty"`
	Hour         int    `json:"hour,omitempty"`
	Minute       int    `json:"minute,omitempty"`
	Seconds      int    `json:"seconds,omitempty"`
	MilliSeconds int    `json:"milliSeconds,omitempty"`
	DateTime     string `json:"dateTime,omitempty"`
	Date         string `json:"date,omitempty"`
	Time         string `json:"time,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
	DayOfWeek    string `json:"dayOfWeek,omitempty"`
	DstActive    bool   `json:"dstActive,omitempty"`
}

type ErrorResponse struct {
	Type    string              `json:"type"`
	Title   string              `json:"title"`
	Status  int                 `json:"status"`
	TraceID string              `json:"traceId"`
	Errors  map[string][]string `json:"errors"`
}
