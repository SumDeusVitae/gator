package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SumDeusVitae/gator/internal/config"
	"github.com/SumDeusVitae/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	// Create user
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	})
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			fmt.Println("User with that name already exists")
			os.Exit(1)
		}
		return fmt.Errorf("error creating user: %w", err)
	}

	// update user in config
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User created successfully: %v\n", user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	err = s.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}
func handlerReset(s *state, cmd command) error {

	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users table: %w", err)
	}
	err = s.Config.SetUser("Unknown")
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("Reset successful")
	return nil
}

func handleUsers(s *state, cmd command) error {
	// Read cfg
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("error reading config: %v", err)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}
	for _, user := range users {
		if user == cfg.CurrentUserName {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Println(user)
		}

	}

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
