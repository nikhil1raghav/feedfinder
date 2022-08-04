package clients

import (
	"log"
	"net/url"
	"strings"

	"github.com/nikhil1raghav/feedfinder/values"

	"github.com/gocolly/colly"
)

type Crawler struct {
	*colly.Collector
}

func (c *Crawler) TypePass(u string) []string {
	links := make([]string, 0)
	base, err := url.Parse(u)
	if err != nil {
		log.Println("Error parsing url", err)
	}
	c.OnHTML("link[type]", func(e *colly.HTMLElement) {
		for _, linkType := range values.FeedTypes {
			if e.Attr("type") == linkType {
				feedUrl, _ := url.Parse(e.Attr("href"))
				links = append(links, base.ResolveReference(feedUrl).String())
			}
		}
	})
	c.Visit(u)
	log.Printf("Found %d feed links in typePass", len(links))
	return links
}
func (c *Crawler) GetAllAnchors(u string) []string {

	anchors := make([]string, 0)
	base, err := url.Parse(u)
	if err != nil {
		log.Println("Error parsing url", err)
		return []string{}
	}
	log.Println("Getting all anchors")

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Count(e.Attr("href"), "://") > 0 {
			anchors = append(anchors, e.Attr("href"))
		} else {
			log.Println("Got this local url", e.Attr("href"))
			localUrl, err := url.Parse(e.Attr("href"))
			if err != nil {
				log.Println("error parsing local url", err)
			} else {

				anchors = append(anchors, base.ResolveReference(localUrl).String())
				log.Println("Appended", base.ResolveReference(localUrl).String())
			}
		}
	})
	c.Visit(u)
	return anchors
}
