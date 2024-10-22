package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Println("Failed to make reqeust")
		return &RSSFeed{}, err
	}

	// ...
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req.Header.Add("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to get response")
		return &RSSFeed{}, err
	}
	// Checking response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected response status: %v\n", resp.StatusCode)
		return &RSSFeed{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response")
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		fmt.Println("Failed to unmarshall xml")
		return &RSSFeed{}, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}
	return &rssFeed, nil
}
