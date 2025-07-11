package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/utkarshjagtap/gator/internal/config"
	"github.com/utkarshjagtap/gator/internal/database"
)

func main() {

	configf, err := config.Read()
	if err != nil {
		fmt.Println("There was an error reading conf", err)
		return
	}

	db, err := sql.Open("postgres", configf.Db_url)
	dbQueries := database.New(db)

	stateU := state{
		config: &configf,
		db:     dbQueries,
	}

	commands := commands{
		commandsmap: make(map[string]func(s *state, cmd command) error),
	}

	commands.register("login", handlerLogins)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handleAgg)
	commands.register("addfeed", middlewearLoggedIn(handleAddFeed))
	commands.register("feeds", handleFeeds)
	commands.register("follow", middlewearLoggedIn( handleFollow))
  commands.register("following", middlewearLoggedIn(handleFollowing))
  commands.register("unfollow", middlewearLoggedIn(handleUnfollow))
  commands.register("browse", middlewearLoggedIn(handlerBrowse))
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("invalid arguments", arguments)
		os.Exit(1)
	}

	err = commands.run(&stateU, command{
		name:      arguments[1],
		arguments: arguments[2:],
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func middlewearLoggedIn(handle func(s *state, cmd command, user database.User)error) func(*state, command)error{
  return func (s *state, cmd command) error {
    user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
    if err != nil{
      return err
    }
    return handle(s, cmd, user)
  }
}



