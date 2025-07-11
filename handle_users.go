package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return fmt.Errorf("invalid usage of command %s", cmd.name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) == 0 {
		fmt.Println("no users found")
		return nil
	}
	current_user := s.config.Current_user_name

	for _, user := range users {
		if user == current_user {
			fmt.Println("*", user, "(current)")
			continue
		}
		fmt.Println("*", user)
	}
	return nil
}
