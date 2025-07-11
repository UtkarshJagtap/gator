package main

import (
	"context"
	"fmt"

	"github.com/utkarshjagtap/gator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error{

  if len(cmd.arguments) > 1{
    return fmt.Errorf("invalid usage of %s", cmd.name)
  }

  feedid, err := s.db.GetFeedIdByURL(context.Background(), cmd.arguments[0]) 
  if err != nil{
    return fmt.Errorf("unable to find feed for given url %v", err)
  }

  err = s.db.DeleteFeedFollowForUser(context.Background(), database.DeleteFeedFollowForUserParams{
    UserID : user.ID,
    FeedID: feedid,
  })

  if err != nil {
    return fmt.Errorf("There was an error while unfollowing %v", err)
  }

  fmt.Println("Unfollowed")




  return nil
}
