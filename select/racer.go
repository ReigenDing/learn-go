package racer

import (
	"net/http"
	"time"
)

func websiteRacer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	bDuration := time.Since(startB)
	if aDuration < bDuration {
		return b
	}

	return a
}
