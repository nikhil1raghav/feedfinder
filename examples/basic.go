package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/nikhil1raghav/feedfinder"
	"github.com/nikhil1raghav/feedfinder/values"
)

func main() {

	f := feedfinder.NewFeedFinder(
		feedfinder.UserAgent(values.ChromeUserAgent),
		feedfinder.CheckAll(true),
		feedfinder.TimeOut(10*time.Second),
	)
	url := flag.String("URL", "https://raghavnikhil.com/", "url to find feeds for")
	flag.Parse()
	links, _ := f.FindFeeds(*url)
	for _, link := range links {
		fmt.Println(link)
	}
}
