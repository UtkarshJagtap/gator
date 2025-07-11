package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/utkarshjagtap/gator/internal/database"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("invalid command %s", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("invalid duration %s, %v", cmd.arguments[0], err)
	}

	fmt.Println("Collecting feeds for every", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {

	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("could not get next feed to fetch")
		return
	}

	scrapeFeed(s.db, next)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	//mark it as fetched

	_, err := db.MarKFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("There was an error marking the feed", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println("Error while fetching feed", feed.Name)
		return
	}



	for _, items := range feedData.Channel.Item {
  publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, items.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		err = db.CreatePost(context.Background(), database.CreatePostParams{
      ID: uuid.New(),
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      FeedID: feed.ID,
      Title: items.Title,
      Url: items.Link,
      Description: sql.NullString{
      },
      PublishedAt: publishedAt,
    },)
    
		fmt.Println("found new post", items.Title)
	}

	log.Printf("Feed %s colllected %v posts found", feed.Name, len(feedData.Channel.Item))

}
