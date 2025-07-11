package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/utkarshjagtap/gator/internal/database"
)
func handlerRegister(s *state, cmd command) error {

	if len(cmd.arguments) == 0 || len(cmd.arguments) > 1 {
		return fmt.Errorf("invalid command %s", cmd.name)
	}

	//checking if the user is already there
	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])

	//the function will return an error if the user is not there
	if err == sql.ErrNoRows {

		parm := database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.arguments[0],
		}

		user, err = s.db.CreateUser(context.Background(), parm)
		if err != nil {
			return fmt.Errorf("%s there was error registering, %v", cmd.arguments[0], err)
		}

		err := s.config.SetUser(user.Name)
		if err != nil {
			return err
		}

		fmt.Println("Registered as", user.Name)
		return nil

	}

	if user.Name != "" {
		return fmt.Errorf("%s already exist", user.Name)
	}

	return err
}
