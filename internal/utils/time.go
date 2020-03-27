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
	parsedExpirationTime, err := time.Parse("2006-01-02T15:04:05-0700", dateTime)
	if err != nil {
		log.Fatal(err)
	}
	return parsedExpirationTime.Before(time.Now())
}
