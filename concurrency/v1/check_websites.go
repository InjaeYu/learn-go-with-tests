package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		result[url] = wc(url)
	}

	return result
}
