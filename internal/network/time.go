package network

import "time"

const isoFormat = "2006-01-02T15:04:05-0700"

func CalculateExpirationDate(expiresIn int) string {
	now := time.Now()
	expirationDuration := time.Duration(expiresIn) * time.Second
	return now.Add(expirationDuration).Format(isoFormat)
}
