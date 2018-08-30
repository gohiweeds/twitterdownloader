package twitterdownloader

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	TweetVideoUrl        = "og:video:url"
	TweetVideoSecure_url = "og:video:secure_url"
)

// Video URL: <meta  property="og:video:url" content="https://twitter.com/i/videos/1033716646911729665?embed_source=facebook">
func (c *Client) ParseVideoUrl(reader io.ReadCloser) string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	var videoUrl string
	// Find the review items
	doc.Find("meta").Each(func(index int, element *goquery.Selection) {
		property, exist := element.Attr("property")
		if exist {
			// fmt.Println("property:", property)
			if 0 == strings.Compare(property, TweetVideoUrl) ||
				0 == strings.Compare(property, TweetVideoSecure_url) {
				var exists bool
				videoUrl, exists = element.Attr("content")
				if exists {
					return
				}
			}
		}
	})
	return videoUrl
}

func (c *Client) ParseVideoJS(reader io.ReadCloser) string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	var jsUrl string
	doc.Find("script").Each(func(index int, element *goquery.Selection) {
		js, exist := element.Attr("src")
		if exist {
			jsUrl = js
		}
	})
	return jsUrl
}
