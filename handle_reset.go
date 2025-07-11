package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error{

  err := s.db.DeletUsers(context.Background())
  if err != nil{
    return err
  }

  fmt.Println("Deleted all the users")

  return nil
} 
