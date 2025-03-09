package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func ConvertUnixTimeToObjectTime(t string) (time.Time, error) {
	// Convert the string to an integer
	unixTimestamp, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing Unix timestamp %s \n", err.Error())
		return time.Time{}, errors.New("error convert unix time")
	}
	// Convert Unix timestamp to time.Time
	timestamp := time.Unix(unixTimestamp, 0)
	return timestamp, nil
}

func SetThailandTimezone(t time.Time) time.Time {
	thailandLocation, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading Thailand timezone:", err)
		return t
	}

	t = t.In(thailandLocation)
	// t = time.Date(
	// 	t.Year(),
	// 	t.Month(),
	// 	t.Day(),
	// 	t.Hour(),
	// 	t.Minute(),
	// 	t.Second(),
	// 	t.Nanosecond(),
	// 	thailandLocation,
	// )

	return t
}
