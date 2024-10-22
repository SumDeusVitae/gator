package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/SumDeusVitae/gator/internal/config"
	"github.com/SumDeusVitae/gator/internal/database"
)

type state struct {
	Config *config.Config
	db     *database.Queries
}

func main() {

	// Load config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}

	// Loading DB
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close() // Don't forget to close the connection when you're done

	dbQueries := database.New(db)

	currentState := &state{
		Config: &cfg,
		db:     dbQueries,
	}

	currentCommands := commands{
		Handlers: make(map[string]func(*state, command) error),
	}

	// command handlers
	currentCommands.register("login", handlerLogin)
	currentCommands.register("read", handlerReadConfig)
	currentCommands.register("register", handlerRegister)
	currentCommands.register("reset", handlerReset)
	currentCommands.register("users", handlerUsers)
	currentCommands.register("agg", handlerAgg)
	currentCommands.register("addfeed", handlerAddFeed)
	currentCommands.register("feeds", handlerAllFeeds)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = currentCommands.run(currentState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
