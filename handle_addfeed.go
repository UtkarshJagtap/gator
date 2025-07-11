package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/utkarshjagtap/gator/internal/database"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("invalid usage")
	}

	current_user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("unable to fetch current user %v", err)
	}

	newfeed, err := s.db.CreateNewFeed(context.Background(), database.CreateNewFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    current_user.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create new feed %v", err)
	}

	fmt.Println(newfeed.Name)
	fmt.Println(newfeed.Url)

	followed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    current_user.ID,
		FeedID:    newfeed.ID,
	},
	)

	fmt.Println(current_user.Name, "has added", newfeed.Name)
	fmt.Println(followed.UserName, "has followed", followed.FeedName)

	return nil
}
