package concurrency

type WebsiteCheck func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteCheck, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		// result := wc(url)
		go func(u string) {

			resultsChannel <- result{u, wc(u)}
		}(url)

	}
	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.string] = result.bool
	}
	return results
}
