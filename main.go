package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

func main() {
	log.Println("Rss Tool")

	feedLoc := "https://bigdatahebdo.com/podcast/index.xml"
	log.Println("Fechting feed : ", feedLoc)
	feedData, err := fetchFeed(feedLoc)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseString(feedData)
	log.Println("Title : ", feed.Title)
	log.Printf("Updated : %s (%s)", feed.UpdatedParsed, feed.Updated)

	for i := range feed.Items {
		item := feed.Items[i]
		//          2022/12/18 20:41:53
		fmt.Printf("                   %+v\n", item.Title)
		fmt.Printf("                   %+v\n", item.Published)
		fmt.Printf("                   %+v\n", item.Authors)
		fmt.Printf("                   %+v\n", item.Categories)
		//fmt.Printf("                   %+v\n", item.Description)
		fmt.Printf("                   %+v\n", item.Link)
		fmt.Printf("\n")
	}

	//fmt.Printf("%+v", feed)

}

func fetchFeed(feedLoc string) (string, error) {
	if strings.HasPrefix(feedLoc, "http") {
		return fetchURL(feedLoc)
	}
	file, err := fetchFile(feedLoc)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func fetchFile(path string) (string, error) {
	f, err := ioutil.ReadFile(path)
	return string(f), err
}

func fetchURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}
