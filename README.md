# Feedfinder

Golang port of [this](https://github.com/dfm/feedfinder2) python library (not actively maintained).


It finds links to rss/atom/rdf feeds in a website. 
Better support for Twitter and reddit links and option to add more extensions

Wrote this to use in rewrite of this [bot](https://github.com/nikhil1raghav/rssbot) in golang.

```go
f:=feedfinder.NewFeedFinder()
url:="old.reddit.com/r/unixporn"
feeds, _:=f.FindFeeds(url)
for _, feed:=range feeds{
    fmt.Println(feed)
}
```
