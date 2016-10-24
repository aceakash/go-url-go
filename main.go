package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	urls := []string {
		"https://github.com/golang",
		"https://github.com/python",
	}
	fmt.Println(fetchUrls(urls))
}

type urlAndRespBody struct {
	url      string
	respBody string
}

func fetchUrls(urls []string) ([]urlAndRespBody, error) {
	uAndRs := make([]urlAndRespBody, 0)
	for _, u := range urls {
		r, err := fetchUrl(u)
		if err != nil {
			continue  // todo handle
		}
		uAndR := urlAndRespBody{
			url: u,
			respBody: r,
		}
		uAndRs = append(uAndRs, uAndR)
	}
	return uAndRs, nil
}

func fetchUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func Concatenator(url ...string) (megabody string, err error) {
	urlsAndBodies, _ := fetchUrls(url)
	for _, u := range urlsAndBodies {
		megabody += u.respBody
	}
	return
}
