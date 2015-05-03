package loader

import (
  r "github.com/jteeuwen/go-pkg-rss"
)

func itemHandler(feed *r.Feed, ch *r.Channel, newitems []*r.Item) {
	for _, item := range newitems[0:5] {
		pp(" - " + item.Title + "\n")
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
