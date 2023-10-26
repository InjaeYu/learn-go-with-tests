package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			result[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return result
}
