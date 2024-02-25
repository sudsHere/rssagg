package main

import (
	"context"
	"database/sql"
	"github/sudsHere/rssagg/internal/database"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on %v goRoutines with %s timeIntervals", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetFeedsForFetching(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feed: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(wg, db, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Println("Error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeeds(feed.Url)

	if err != nil {
		log.Println("Error fetching feed: ", err)
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("couldn't parse date %v with error %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubTime,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post:", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
