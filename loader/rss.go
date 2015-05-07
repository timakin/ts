package loader

import (
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"strings"
	r "github.com/jteeuwen/go-pkg-rss"
)

type RSS struct {
    XMLName xml.Name `xml:"rss"`
    Items Items `xml:"channel"`
}

type Items struct {
    XMLName xml.Name `xml:"channel"`
    ItemList []Item `xml:"item"`
}

type Item struct {
    Title string `xml:"title"`
    Link string `xml:"link"`
		Description string `xml:"description"`
}

func getXMLDataFromUrl(url string) (xmlRes RSS) {
	resp, err := http.Get(url)
	perror(err)
	xmlDataFromHttp, err := ioutil.ReadAll(resp.Body)
	perror(err)
	err = xml.Unmarshal([]byte(xmlDataFromHttp), &xmlRes)
	perror(err)
	return xmlRes
}

func GetRssFeed(name string, url string, rss chan ResultData) {
	var result ResultData
	var RSSTitle, RSSUrl []string

	xml := getXMLDataFromUrl(url)
	for _, item := range xml.Items.ItemList[0:5] {
		RSSTitle = append(RSSTitle, strings.Replace(item.Title,"\n","",-1))
		RSSUrl = append(RSSUrl, strings.Replace(item.Link,"\n","",-1))
	}
	result.Setter(name, RSSTitle, RSSUrl)
	rss <- result
}

func GetRssFeedWithDesc(name string, url string, rss chan ResultData) {
	var result ResultData
	var RSSTitle, RSSUrl []string

	xml := getXMLDataFromUrl(url)
	for _, item := range xml.Items.ItemList[0:5] {
		RSSTitle = append(RSSTitle, strings.Replace(item.Title + ": " + removeBreak(item.Description),"\n","",-1))
		RSSUrl = append(RSSUrl, strings.Replace(item.Link,"\n","",-1))
	}
	result.Setter(name, RSSTitle, RSSUrl)
	rss <- result
}


func itemHandler(feed *r.Feed, ch *r.Channel, newitems []*r.Item) {
	for _, item := range newitems[0:5] {
		pp(" - " + item.Title + "\n")
		pp("   - " + item.Links[0].Href + "\n")
	}
	pp("\n")
}

func itemHandlerWithDescription(feed *r.Feed, ch *r.Channel, newitems []*r.Item) {
	for _, item := range newitems[0:5] {
		pp(" - " + item.Title + ": " + removeBreak(item.Description) + "\n")
		pp("   - " + item.Links[0].Href + "\n")
	}
	pp("\n")
}

func GetRSSFeed(uri string) {
	timeout := 5
	feed := r.New(timeout, true, nil, itemHandler)
	err := feed.Fetch(uri, nil)
	perror(err)
}

func GetRSSFeedWithDescription(uri string) {
	timeout := 5
	feed := r.New(timeout, true, nil, itemHandlerWithDescription)
	err := feed.Fetch(uri, nil)
	perror(err)
}
