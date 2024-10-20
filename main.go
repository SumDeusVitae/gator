package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SumDeusVitae/gator/internal/config"
)

type state struct {
	Config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currentState := &state{Config: &cfg}

	currentCommands := commands{
		Handlers: make(map[string]func(*state, command) error),
	}

	// command handlers
	currentCommands.register("login", handlerLogin)
	currentCommands.register("read", handlerReadConfig)

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
