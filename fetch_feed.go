package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
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
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("There was an error creating new request %v", err)
	}

	request.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("There was an error making a new request %v", err)
	}

	if res.StatusCode > 199 && res.StatusCode < 300 {
		reader := io.Reader(res.Body)
		defer res.Body.Close()

		data, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("There was an error reading the response %v", err)
		}

		var rssfeed RSSFeed
		err = xml.Unmarshal(data, &rssfeed)
		if err != nil {
			return nil, fmt.Errorf("There was an erro marshalling the response %v", err)
		}

    rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
    rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)

    for _, items := range rssfeed.Channel.Item{
    items.Title = html.UnescapeString(items.Title)
    items.Description = html.UnescapeString(items.Description)
    }
    

		return &rssfeed, nil

	}
	return nil, fmt.Errorf("Request failed with response code %d", res.StatusCode)
}
