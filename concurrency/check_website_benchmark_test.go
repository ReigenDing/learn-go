package concurrency

import (
	"testing"
	"time"
)

func SlowStubWebsiteChecker(_ string) bool {
	time.Sleep(2 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	println("===========")
	for i := 0; i < b.N; i++ {
		CheckWebsite(SlowStubWebsiteChecker, urls)
	}
}
