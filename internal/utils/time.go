package utils

import (
	"log"
	"time"
)

const isoFormat = "2006-01-02T15:04:05-0700"

func CalculateExpirationDate(expiresIn int) string {
	now := time.Now()
	expirationDuration := time.Duration(expiresIn) * time.Second
	return now.Add(expirationDuration).Format(isoFormat)
}

func IsAfter(dateTime string) bool {
	parsedExpirationTime, err := time.Parse(isoFormat, dateTime)
	if err != nil {
		log.Fatal(err)
	}
	return parsedExpirationTime.Before(time.Now())
}

func ParseDateTime(dateTime string) *time.Time {
	result, err := time.Parse("2006-01-02T15:04:05Z", dateTime)

	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func ParseDate(date string, precision string) *time.Time {
	if precision == "day" {
		return parseDate(date, "2006-01-02")
	}
	return parseDate(date, "2006")
}

func parseDate(date string, layout string) *time.Time {
	result, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}
