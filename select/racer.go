package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := MeasureResponseTime(a)
	durationB := MeasureResponseTime(b)

	if durationA < durationB {
		return a
	}

	return b
}

func MeasureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}
