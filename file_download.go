package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	doc, err := goquery.NewDocument("URL")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		download(url)

	})
	t = time.Now()
	fmt.Println(t)
}

func download(url string) (err error) {
	filename := path.Base(url)
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}
