package mongo

import (
	"log"
	"time"
)

func convertToThailandTime(t time.Time) time.Time {
	thailandLoc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("WARNING: Failed to load Thailand time zone: %v. Using system time zone.", err)
		return t
	}
	return t.In(thailandLoc)
}
