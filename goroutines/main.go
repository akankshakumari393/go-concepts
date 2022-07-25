package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// go to www.makemytrip.com/ programmatically and download the content of mmt.com. Identify the number of occurences of makemytrip keyword.

func main() {
	var wg = &sync.WaitGroup{}
	urls := []string{"https://www.makemytrip.com/", "https://www.goibibo.com/"}
	keywords := []string{"makemytrip", "goibibo.com"}
	for i, url := range urls {
		wg.Add(1)
		go CountOccurence(url, keywords[i], wg)
	}

	wg.Wait()
}

func CountOccurence(url, keyword string, wg *sync.WaitGroup) {
	defer wg.Done()
	content, err := downloadContent(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	c := CountOccurences(content, keyword)
	fmt.Println("count of", keyword, url, c)
}

// download the content from a given URL
func downloadContent(url string) (string, error) {
	resp, _ := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func CountOccurences(input string, keyword string) int {
	return strings.Count(input, keyword)
}
