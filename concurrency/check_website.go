package concurrency

type WebsiteCheck func(string) bool

func CheckWebsite(wc WebsiteCheck, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		// result := wc(url)
		results[url] = wc(url)

	}

	return results
}
