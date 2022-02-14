package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	urls := []string{}
	for i := 0; i < 10; i++ {
		urls = append(urls, fmt.Sprintf("http://test.com/%d", i))
	}
	success := getURLs(urls)
	log.Println("program has been completed with:", success)
}

func getURLs(urls []string) string {
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			log.Println("get url", i, ":", url)
			getOneURL(url) // ignore error for this task
		}(i, url)
	}

	wg.Wait()
	return "success"
}

func getOneURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(ioutil.ReadAll(resp.Body))
	return nil
}
