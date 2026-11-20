package medium

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"

	"portfolio/internal/models"
)

// Sync fetches the feed once and upserts every post into the articles table,
// keyed on external_url. It returns how many posts were in the feed.
func Sync(ctx context.Context, db *gorm.DB, client *http.Client, username string) (int, error) {
	posts, err := Fetch(ctx, client, username)
	if err != nil {
		return 0, err
	}

	for _, p := range posts {
		fields := models.Article{
			Title:       p.Title,
			Summary:     p.Summary,
			ImageURL:    p.ImageURL,
			Tags:        strings.Join(p.Tags, ","),
			Platform:    "medium",
			PublishedAt: p.PublishedAt,
		}
		// FirstOrCreate finds the row by URL (or creates it); Assign makes
		// sure an existing row is updated with the latest title/summary/etc.
		var a models.Article
		if err := db.WithContext(ctx).
			Where(models.Article{ExternalURL: p.URL}).
			Assign(fields).
			FirstOrCreate(&a).Error; err != nil {
			return 0, err
		}
	}
	return len(posts), nil
}

// StartSync runs Sync immediately and then every interval until ctx is
// cancelled. It runs in its own goroutine and never blocks server start-up;
// failures are logged and retried on the next tick.
func StartSync(ctx context.Context, db *gorm.DB, username string, interval time.Duration) {
	client := &http.Client{Timeout: 15 * time.Second}

	run := func() {
		n, err := Sync(ctx, db, client, username)
		if err != nil {
			log.Printf("medium: sync failed: %v", err)
			return
		}
		log.Printf("medium: synced %d posts from @%s", n, strings.TrimPrefix(username, "@"))
	}

	go func() {
		run()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				run()
			}
		}
	}()
}
