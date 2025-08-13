package rss

// import (
// 	"context"
// 	"encoding/xml"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type RSSFeed struct {
// 	Channel struct {
// 		Title       string    `xml:"title"`
// 		Link        string    `xml:"link"`
// 		Description string    `xml:"description"`
// 		Item        []RSSItem `xml:"item"`
// 	} `xml:"channel"`
// }

// type RSSItem struct {
// 	Title       string `xml:"title"`
// 	Link        string `xml:"link"`
// 	Description string `xml:"description"`
// 	PubDate     string `xml:"pubDate"`
// }

// func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
// 	// init request
// 	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
// 	if err != nil {
// 		return &RSSFeed{}, err
// 	}

// 	// set headers
// 	req.Header.Set("User-Agent", "gator")

// 	client := &http.Client{}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return &RSSFeed{}, err
// 	}

// 	defer res.Body.Close()

// 	// transform data
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return &RSSFeed{}, err
// 	}
// 	var feed RSSFeed
// 	err = xml.Unmarshal(data, &feed)
// 	if err != nil {
// 		return &feed, err
// 	}
// 	fmt.Println("FEED is: ", feed)
// 	return &feed, nil

// }
