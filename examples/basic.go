package main

import (
	"fmt"

	"github.com/nikhil1raghav/feedfinder"
)

func main() {
	f := feedfinder.NewFeedFinder(feedfinder.UserAgent("Some useragent"))
	url := "old.reddit.com/r/unixporn/"
	links, _ := f.FindFeeds(url)
	for _, link := range links {
		fmt.Println(link)
	}
}
