package main

import (
	"fmt"

	"github.com/utkarshjagtap/gator/internal/config"
	"github.com/utkarshjagtap/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandsmap map[string]func(s *state, cmd command) error
}

func (cmnds commands) run(s *state, cmd command) error {
	fu, ok := cmnds.commandsmap[cmd.name]
  if !ok{
    return fmt.Errorf("command not found %s", cmd.name)
  }
	err := fu(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (cmnds commands) register(name string, f func(*state, command) error) {
	cmnds.commandsmap[name] = f
}




