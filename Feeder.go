package feedfinder

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/nikhil1raghav/feedfinder/clients"
	"github.com/nikhil1raghav/feedfinder/utils"
	"github.com/nikhil1raghav/feedfinder/values"

	"github.com/gocolly/colly"
)

type FeedFinder struct {
	UserAgent  string
	TimeOut    time.Duration
	CheckAll   bool
	HttpClient http.Client
	Crawler    clients.Crawler
}

func NewFeedFinder(options ...func(*FeedFinder)) *FeedFinder {
	f := &FeedFinder{}
	f.Init()

	return f
}
func UserAgent(ua string) func(*FeedFinder) {
	return func(f *FeedFinder) {
		f.UserAgent = ua
		f.Crawler.UserAgent = ua
	}
}
func CheckAll(checkall bool) func(*FeedFinder) {
	return func(f *FeedFinder) {
		f.CheckAll = checkall
	}
}
func TimeOut(timeout time.Duration) func(*FeedFinder) {
	return func(f *FeedFinder) {
		f.TimeOut = timeout
	}
}
func (f *FeedFinder) Init() {
	f.UserAgent = values.ChromeUserAgent
	f.CheckAll = false
	f.TimeOut = 60
	f.Crawler.Collector = colly.NewCollector(colly.UserAgent(f.UserAgent))
}
func (f *FeedFinder) getFeed(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", f.UserAgent)
	resp, err := f.HttpClient.Get(url)
	if err != nil {
		log.Printf("Error getting %s, %s", url, err.Error())
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (f *FeedFinder) isFeedData(data []byte) bool {
	dataString := strings.ToLower(string(data))
	if strings.Count(dataString, "<html") > 0 {
		return false
	}
	for _, header := range values.FeedHeaders {
		if strings.Count(dataString, header) > 0 {
			return true
		}
	}
	return false
}
func (f *FeedFinder) isFeedUrl(url string) bool {
	url = strings.ToLower(url)
	for _, suffix := range values.FeedUrlSuffix {
		if strings.HasSuffix(url, suffix) {
			return true
		}
	}
	return false
}
func (f *FeedFinder) isFeedLike(url string) bool {
	url = strings.ToLower(url)
	for _, word := range values.FeedLike {
		if strings.Count(url, word) > 0 {
			return true
		}
	}
	return false
}

//called when all else failed
//validate feed after guessing
func (f *FeedFinder) guessUrls(u string) []string {
	guessed := make([]string, 0)
	for _, suffix := range values.GuessWords {
		url, err := utils.JoinUrl(u, suffix)
		if err != nil {
			log.Println(err)
			continue
		} else if validFeed, _ := f.isFeed(url); validFeed {
			guessed = append(guessed, url)
		}
	}
	return guessed

}
func (f *FeedFinder) isFeed(u string) (bool, error) {
	data, err := f.getFeed(u)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return f.isFeedData(data), nil
}
func (f *FeedFinder) FindFeeds(url string) ([]string, error) {
	url = utils.ForceUrl(url)
	feedUrls := make([]string, 0)
	if validFeed, _ := f.isFeed(url); validFeed {
		feedUrls = append(feedUrls, url)
	}

	feedUrls = append(feedUrls, f.Crawler.TypePass(url)...)

	if len(feedUrls) > 0 && !f.CheckAll {
		return feedUrls, nil
	}

	anchors := f.Crawler.GetAllAnchors(url)
	filteredUrl := make([]string, 0)
	for _, anchor := range anchors {
		if f.isFeedUrl(anchor) {
			filteredUrl = append(filteredUrl, anchor)
		} else if f.isFeedLike(anchor) {
			filteredUrl = append(filteredUrl, anchor)
		}
	}
	for _, u := range filteredUrl {
		if validFeed, _ := f.isFeed(u); validFeed {
			feedUrls = append(feedUrls, u)
		}
	}

	if len(feedUrls) > 0 && !f.CheckAll {
		return feedUrls, nil
	}

	feedUrls = append(feedUrls, f.guessUrls(url)...)

	return feedUrls, nil
}
