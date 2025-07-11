package main

import (
	"context"
	"fmt"
)

func handleFeeds(s *state, cmd command) error {
  if len(cmd.arguments) > 0{
    return fmt.Errorf("invalid usage %s", cmd.name)
  }

  feeds, err := s.db.GetFeeeds(context.Background())
  if err != nil{
    return fmt.Errorf("unable to list the feeds %v", err)
  }

  for _, feed := range feeds{
    fmt.Println(feed.Name, feed.Url, feed.UserName)
  }

	return nil
}
