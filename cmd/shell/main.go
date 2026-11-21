package main

import (
	"fmt"
	"go-fiber-jwt/database"
	"go-fiber-jwt/models"
	"os"

	"github.com/c-bata/go-prompt"
)

func executor(in string) {
	switch in {
	case "exit":
		fmt.Println("Exiting shell...")
		os.Exit(0)
	case "list-users":
		var users []models.User
		database.DB.Find(&users)
		fmt.Println("Users:")
		for _, user := range users {
			fmt.Printf("- %s (%s)\n", user.Name, user.Email)
		}
	default:
		fmt.Println("Unknown command")
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "exit", Description: "Exit the shell"},
		{Text: "list-users", Description: "List all users"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	database.ConnectDB()
	defer func() {
		sqlDB, _ := database.DB.DB()
		sqlDB.Close()
	}()

	fmt.Println("Welcome to the application shell!")
	fmt.Println("Available commands:")
	fmt.Println("- exit: Exit the shell")
	fmt.Println("- list-users: List all users")

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("app-shell"),
	)
	p.Run()
}
