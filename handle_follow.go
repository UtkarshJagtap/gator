package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/utkarshjagtap/gator/internal/database"
)

func handleFollow(s *state, cmd command) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("invalid usage: %s", cmd.name)
	}

	feed_id, err := s.db.GetFeedIdByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("There was an error finding the feed for given url %s, %v", cmd.arguments[0], err)
	}

	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("There was an error finding the current user %s, %v", s.config.Current_user_name, err)
	}

	followed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed_id,
	},
	)

	fmt.Println(followed.UserName, "has followed", followed.FeedName)

	return nil
}
