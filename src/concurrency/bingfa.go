package concurrency

type WebsiteChecker = func(url string) bool

func CheckWebsites(wc WebsiteChecker, websites []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range websites {
		results[url] = wc(url)
	}
	return results
}
