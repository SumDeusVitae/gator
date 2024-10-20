package main

import (
	"fmt"

	"github.com/SumDeusVitae/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	err := s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerReadConfig(s *state, cmd command) error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("error reading config: %v", err)
	}
	// Print the current configuration details.
	fmt.Printf("Current Configuration:\n")
	fmt.Printf("Database URL: %s\n", cfg.DbURL)
	fmt.Printf("Current User Name: %s\n", cfg.CurrentUserName)

	return nil
}
