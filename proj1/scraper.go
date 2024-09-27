package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/SandeepSinghSethi/mygoproj/internal/database"
	"github.com/google/uuid"
)

func scrapeRss(db *database.Queries, concurrency int, timeBtwRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s seconds !!", concurrency, timeBtwRequest)
	ticker := time.NewTicker(timeBtwRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds :", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait() // waiting until concurrency times ; wg is a queue like it is for threading.Wait in python
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error while updating timestamps :", err)
	}

	rssfeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error while fetching url :", err)
		return
	}

	for _, item := range rssfeed.Channel.Item {
		// log.Println("Found Post ", item.Title, " on feed ", feed.Name)
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pdate, err := time.Parse(time.RFC822, item.PubDate)
		if err != nil {
			pdate, err = time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				pdate, err = time.Parse(time.RFC1123, item.PubDate)
				if err != nil {
					log.Printf("couldn't parse date : %v with error : %v", item.PubDate, err)
					continue
				}
			}
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			ModifiedAt:  time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pdate,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("couldnt create post : ", err)
		}
	}

	log.Printf("Feed %s collected with %v posts found ", feed.Name, len(rssfeed.Channel.Item))
}
