package clients

import (
	"fmt"
	"github.com/nikhil1raghav/feedfinder/values"
	"testing"
	"time"
)

func TestUserAgent(t *testing.T){
	c:=NewCrawler(values.ChromeUserAgent, time.Minute)
	resp,err:=c.Get("http://localhost:8080")
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(resp)
}
