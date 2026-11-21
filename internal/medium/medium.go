// Package medium reads a Medium profile's public RSS feed and keeps the
// articles table in sync with it.
//
// Medium exposes every profile at https://medium.com/feed/@<username>. The
// feed only ever contains the 10 most recent stories, so syncing adds and
// updates articles but never deletes older ones that have dropped off it.
package medium

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Post is one story from the feed, already cleaned up for storage.
type Post struct {
	Title       string
	URL         string
	Summary     string
	ImageURL    string
	Tags        []string
	PublishedAt time.Time
}

// rss mirrors only the parts of the RSS document we need.
// `content:encoded` lives in the namespace below; encoding/xml matches it
// with the "namespace-URL local-name" form in the struct tag.
type rss struct {
	Channel struct {
		Items []struct {
			Title      string   `xml:"title"`
			Link       string   `xml:"link"`
			PubDate    string   `xml:"pubDate"`
			Categories []string `xml:"category"`
			Content    string   `xml:"http://purl.org/rss/1.0/modules/content/ encoded"`
		} `xml:"item"`
	} `xml:"channel"`
}

// FeedURL returns the RSS URL for a Medium username (with or without the "@").
func FeedURL(username string) string {
	return "https://medium.com/feed/@" + strings.TrimPrefix(username, "@")
}

// Fetch downloads and parses a user's feed.
func Fetch(ctx context.Context, client *http.Client, username string) ([]Post, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, FeedURL(username), nil)
	if err != nil {
		return nil, err
	}
	// Medium rejects some requests that have no User-Agent.
	req.Header.Set("User-Agent", "portfolio-medium-sync/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("medium: fetch feed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("medium: fetch feed: unexpected status %d", resp.StatusCode)
	}

	// Cap the body size so a misbehaving upstream can't exhaust memory.
	return Parse(io.LimitReader(resp.Body, 5<<20))
}

// Parse turns raw RSS XML into Posts. It is separate from Fetch so it can be
// unit tested without the network.
func Parse(r io.Reader) ([]Post, error) {
	var doc rss
	if err := xml.NewDecoder(r).Decode(&doc); err != nil {
		return nil, fmt.Errorf("medium: parse feed: %w", err)
	}

	posts := make([]Post, 0, len(doc.Channel.Items))
	for _, it := range doc.Channel.Items {
		published, err := time.Parse(time.RFC1123, it.PubDate) // "Thu, 12 Nov 2026 10:26:18 GMT"
		if err != nil {
			published = time.Now()
		}
		posts = append(posts, Post{
			Title:       strings.TrimSpace(it.Title),
			URL:         cleanURL(it.Link),
			Summary:     summarize(it.Content, 180),
			ImageURL:    firstImage(it.Content),
			Tags:        it.Categories,
			PublishedAt: published,
		})
	}
	return posts, nil
}

// cleanURL drops Medium's tracking query string (?source=rss-…) so the same
// story always maps to the same URL — that URL is how sync finds existing rows.
func cleanURL(raw string) string {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return raw
	}
	u.RawQuery = ""
	u.Fragment = ""
	return u.String()
}

var (
	imgRe       = regexp.MustCompile(`<img[^>]+src="([^"]+)"`)
	paragraphRe = regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	tagRe       = regexp.MustCompile(`<[^>]+>`)
	spaceRe     = regexp.MustCompile(`\s+`)
)

// firstImage returns the first <img src> in the post body — Medium puts the
// cover image there.
func firstImage(body string) string {
	if m := imgRe.FindStringSubmatch(body); m != nil {
		return m[1]
	}
	return ""
}

// summarize builds a plain-text excerpt from the first paragraph(s),
// cut at a word boundary.
func summarize(body string, max int) string {
	var text string
	for _, m := range paragraphRe.FindAllStringSubmatch(body, -1) {
		text += " " + m[1]
		if len(text) > max {
			break
		}
	}
	text = tagRe.ReplaceAllString(text, "")
	text = html.UnescapeString(text)
	text = strings.TrimSpace(spaceRe.ReplaceAllString(text, " "))

	runes := []rune(text)
	if len(runes) <= max {
		return text
	}
	cut := string(runes[:max])
	if i := strings.LastIndex(cut, " "); i > max/2 {
		cut = cut[:i]
	}
	return strings.TrimRight(cut, " ,.;:—-") + "…"
}
