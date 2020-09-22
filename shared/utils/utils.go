package utils

import (
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

//CustomValidator ..
type CustomValidator struct {
	Validator *validator.Validate
}

//Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

//DefaultValidator ...
func DefaultValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}


// IndonesianTime is func get Indonesia current time
func IndonesianTime() time.Time {
	return time.Now().UTC().Add(7 * time.Hour)
}

//ParseStringToTime ...
func ParseStringToTime(timeValue string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	t, err := time.Parse(timeLayout, timeValue)
	if err != nil {
		t = time.Now()
	}

	return t
}

//ParseFloat ...
func ParseFloat(value string) float64 {
	valueFloatType, err := strconv.ParseFloat(value, 64)
	if err != nil {
		valueFloatType = 0.0
	}
	return valueFloatType
}


//SecondsToMinutes ...
func SecondsToMinutes(second int64) int64{
	minute := second/60
	return minute
}
