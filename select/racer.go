package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// func websiteRacer(a, b string) string {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)
// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

func websiteRacer(a, b string) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waitting for %s and %s", a, b)
	}
	// var s struct{}
	// s, ok := <-ping(a)
	// fmt.Println(s, ok)
	// s2, ok2 := <-ping(b)
	// fmt.Println(s2, ok2)
	// return b
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(a):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}
