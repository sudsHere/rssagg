package main

import (
	"github/sudsHere/rssagg/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}
	return feeds
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID 	uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowstoFeedFollows(feedFollows database.FeedFollow) FeedFollows {
	return FeedFollows {
		ID: feedFollows.ID,
		CreatedAt: feedFollows.CreatedAt,
		UpdatedAt: feedFollows.UpdatedAt,
		UserID: feedFollows.UserID,
		FeedID: feedFollows.FeedID,
	}
}

func databaseFeedsFollowstoFeedsFollows(dbFeedsFollows []database.FeedFollow) []FeedFollows {
	feedsFollows := []FeedFollows{}
	for _, dbFeedFollows := range dbFeedsFollows {
		feedsFollows = append(feedsFollows, databaseFeedFollowstoFeedFollows(dbFeedFollows))
	}
	return feedsFollows
}