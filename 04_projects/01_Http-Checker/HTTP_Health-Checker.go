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

func checkSite(s *Site, ch chan<- Site) {
	resp, err := http.Get(s.Url)
	if err != nil {
		ch <- Site{Url: s.Url, StatusCode: 500, StatusText: err.Error()}
		return
	} else {
		defer resp.Body.Close()
		s.StatusCode = resp.StatusCode
	}
	ch <- Site{Url: s.Url, StatusCode: s.StatusCode, StatusText: s.StatusText}
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

	ch := make(chan Site)
	for _, url := range urls {
		go checkSite(&Site{Url: url}, ch)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}
