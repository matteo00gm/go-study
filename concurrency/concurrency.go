package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	urlsChannel := make(chan result)

	for _, url := range urls {
		go func() {
			urlsChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-urlsChannel
		results[r.string] = r.bool
	}

	return results
}
