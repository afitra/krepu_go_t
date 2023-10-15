package models

// Parameter struct is represent a data for parameters model
type ParameterTimeZone struct {
	TimeZone string `validate:"required"`
}
