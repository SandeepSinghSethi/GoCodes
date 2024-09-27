package main

import (
	"time"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at`
	ModifiedAt time.Time `json:"modified_at"`
	Username   string    `json:"name"`
	ApiKey     string    `json:"api_key"`
}

type Feed struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	UserID     uuid.UUID `json:"userid"`
}

func dbuserToUser(dbuser database.User) User {
	return User{
		ID:         dbuser.ID,
		CreatedAt:  dbuser.CreatedAt,
		ModifiedAt: dbuser.ModifiedAt,
		Username:   dbuser.Username,
		ApiKey:     dbuser.ApiKey,
	}
}

func dbfeedToFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:         dbfeed.ID,
		CreatedAt:  dbfeed.CreatedAt,
		ModifiedAt: dbfeed.ModifiedAt,
		Name:       dbfeed.Name,
		Url:        dbfeed.Url,
		UserID:     dbfeed.UserID,
	}
}

func dbFeedsslicetoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, dbfeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	UserID     uuid.UUID `json:"userid"`
	FeedID     uuid.UUID `json:"feedid`
}

func dbfeedfollowToFeedFollow(dbfeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:         dbfeedFollow.ID,
		CreatedAt:  dbfeedFollow.CreatedAt,
		ModifiedAt: dbfeedFollow.ModifiedAt,
		UserID:     dbfeedFollow.UserID,
		FeedID:     dbfeedFollow.FeedID,
	}
}

func dbfeedfollowSliceToFeedFollow(dbfeedFollow []database.FeedFollow) []FeedFollow {
	feeds_flw := []FeedFollow{}
	for _, feedFlw := range dbfeedFollow {
		feeds_flw = append(feeds_flw, dbfeedfollowToFeedFollow(feedFlw))
	}
	return feeds_flw
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:published_at`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func dbposttopost(dbpost database.Post) Post {
	var description *string
	if dbpost.Description.Valid {
		description = &dbpost.Description.String
	}
	return Post{
		ID:          dbpost.ID,
		CreatedAt:   dbpost.CreatedAt,
		ModifiedAt:  dbpost.ModifiedAt,
		Title:       dbpost.Title,
		Description: description,
		PublishedAt: dbpost.PublishedAt,
		Url:         dbpost.Url,
		FeedID:      dbpost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbpost := range dbPosts {
		posts = append(posts, dbposttopost(dbpost))
	}
	return posts
}
