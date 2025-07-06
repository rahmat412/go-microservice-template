package date

import (
	"context"
	"time"

	"github.com/rahmat412/go-microservice-template/internal/helper/customerror"
)

func ParseStringToDate(ctx context.Context, dateString string) (time.Time, error) {
	// Parse the date string into a time.Time object
	parseTime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return time.Time{}, customerror.ErrorBirthDateFormat
	}

	return parseTime, nil
}

func ParseDateToString(ctx context.Context, date time.Time) (string, error) {
	// Format the time.Time object into a string
	dateString := date.Format(time.RFC3339)
	if dateString == "" {
		return "", customerror.ErrorBirthDateFormat
	}

	return dateString, nil
}
