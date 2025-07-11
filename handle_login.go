package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogins(s *state, cmd command) error {

	if len(cmd.arguments) == 0 || len(cmd.arguments) > 1 {
		return fmt.Errorf("invalid command %s", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err == sql.ErrNoRows {
		return fmt.Errorf("can not login: %s not found", cmd.arguments[0])
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("logged in as", user.Name)

	return nil
}
