package main

import (
	"fmt"
	"net/http"
)

type Site struct {
	Url        string
	StatusCode int
	StatusText string
}

func checkSite(s *Site) {
	resp, err := http.Get(s.Url)
	if err != nil {
		s.StatusCode = 0
		s.StatusText = "Error"
		return
	} else {
		s.StatusCode = resp.StatusCode
		defer resp.Body.Close()
	}
}
func main() {
	urls := []string{
		"https://github.com",
		"https://cloud.ru",
		"https://music.yandex.ru",
		"https://youtube.com",
		"https://google.com",
		"https://this-domain-does-not-exist-12345.com",
	}
	for _, url := range urls {
		currentSite := Site{Url: url}
		checkSite(&currentSite)
		fmt.Printf("Сайт: %-10s | Код: %-3d | Статус: %s\n",
			currentSite.Url, currentSite.StatusCode, currentSite.StatusText)
	}
}
