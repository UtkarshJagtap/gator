package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error{
  feedUrl := "https://www.wagslane.dev/index.xml"
  if len(cmd.arguments) > 0{
    return fmt.Errorf("invalid command %s", cmd.name)
  }

  feed, err := fetchFeed(context.Background(), feedUrl)
  if err != nil{
    return err
  }

  fmt.Println(feed)
  return nil
}
