package main

import (
	"context"
	"fmt"

	"github.com/utkarshjagtap/gator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error{
  if len(cmd.arguments) > 0{
    return fmt.Errorf("invalid usage for %s", cmd.name)
  }

  following, err := s.db.GetFeedFollowForUser(context.Background(), user.Name)
  if err != nil{
    return fmt.Errorf("there was an error fetching following for %s %v", user.Name, err)
  }

  for _, feed := range following{
    fmt.Println(feed)
  }

  return nil
}
