package main

import (
	"time"

	"github.com/ijt/go-anytime"
	karma "github.com/reconquest/karma-go"
)

func parseDate(date string) (int64, error) {
	var dateUnix int64

	destiny := karma.Describe("method", "parseDate")

	timeNow := time.Now()

	if date == "" {
		dateUnix = timeNow.Unix()
	} else {
		dateParse, err := anytime.Parse(date, timeNow)
		if err != nil {
			return dateUnix, destiny.Describe(
				"error", err,
			).Describe(
				"date", date,
			).Reason(
				"can't convert date to unixtime",
			)
		}
		dateUnix = dateParse.Unix()
	}
	return dateUnix, nil
}

func parseDateTime(date string) (int64, error) {
	timeNow := time.Now()
	dateParse, err := anytime.Parse(date, timeNow)
	if err != nil {
		return 0, karma.Format(err, "can't parse datetime '%s'", date)
	}

	return dateParse.Unix(), nil
}
