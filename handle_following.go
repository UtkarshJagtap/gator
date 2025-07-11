package main

import (
	"context"
	"fmt"
)

func handleFollowing(s *state, cmd command) error{
  if len(cmd.arguments) > 0{
    return fmt.Errorf("invalid usage for %s", cmd.name)
  }

  current_user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
  if err != nil{
    return fmt.Errorf("there was an error fetching current user %s %v", s.config.Current_user_name,err)
  }

  following, err := s.db.GetFeedFollowForUser(context.Background(), current_user.Name)
  if err != nil{
    return fmt.Errorf("there was an error fetching following for %s %v", current_user.Name, err)
  }

  for _, feed := range following{
    fmt.Println(feed)
  }

  return nil
}
