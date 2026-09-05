package rss

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrUnknown = errors.New("unknown error has occurred")
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict = errors.New("conflict")
)

type Feed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title           string `xml:"title"`
	Description     string `xml:"description"`
	Link            string `xml:"link"`
	PublicationDate string `xml:"pubDate"`
}

func Fetch(ctx context.Context, url url.URL) (*Feed, error) {
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", "aggregator")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var feed *Feed
	if err := xml.Unmarshal(bytes, &feed); err != nil {
		return nil, err
	}

	for idx, item := range feed.Channel.Items {
		feed.Channel.Items[idx].Title = html.UnescapeString(item.Title)
		feed.Channel.Items[idx].Description = html.UnescapeString(item.Description)
	}

	return feed, err
}
