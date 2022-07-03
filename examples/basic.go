package main

import (
	"flag"
	"fmt"

	"github.com/nikhil1raghav/feedfinder"
)

func main() {

	f := feedfinder.NewFeedFinder(feedfinder.UserAgent("Some useragent"))
	url := flag.String("URL", "http://lukesmith.xyz", "url to find feeds for")
	flag.Parse()
	links, _ := f.FindFeeds(*url)
	for _, link := range links {
		fmt.Println(link)
	}
}
