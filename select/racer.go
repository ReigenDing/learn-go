package racer

import (
	"fmt"
	"net/http"
	"time"
)

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

func websiteRacer(a, b string) string {
	// select {
	// case <-ping(a):
	// 	return a
	// case <-ping(b):
	// 	return b
	// }
	// var s struct{}
	s, ok := <-ping(a)
	fmt.Println(s, ok)
	s2, ok2 := <-ping(b)
	fmt.Println(s2, ok2)
	return b
}

func ping(url string) chan struct{ a, b string } {
	fmt.Println(url)
	ch := make(chan struct{ a, b string })
	go func() {
		fmt.Println("1")
		http.Get(url)
		fmt.Println("2")
		ch <- struct {
			a string
			b string
		}{"1", "2"}
		close(ch)
		fmt.Println("3")
	}()
	fmt.Println("4")
	s, ok := <-ch
	fmt.Println(s, ok)
	time.Sleep(5 * time.Second)
	return ch
}
