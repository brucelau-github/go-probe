package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	mozillaGet()
}

func mozillaGet() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   60 * time.Second,
	}
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	mozillaHeaders := map[string]string{
		"User-Agent":      `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:58.0) Gecko/20100101 Firefox/58.0`,
		"Accept":          `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`,
		"Accept-Language": `en-US,en;q=0.5`,
		"Accept-Encoding": `gzip, deflate, br`,
		"Cache-Control":   `max-age=0`,
	}
	for k, v := range mozillaHeaders {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Printf("%s", data)
}
