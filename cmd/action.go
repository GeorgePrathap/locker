package cmd

import (
	"fmt"
	"locker/db"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func GetFlagValue(cmd *cobra.Command, flag string) string {
	value, _ := cmd.Flags().GetString(flag)
	return value
}

func AddPassword(cmd *cobra.Command, args []string) {
	user := db.User{
		ID:          uuid.NewString(),
		Title:       GetFlagValue(cmd, "title"),
		Username:    GetFlagValue(cmd, "username"),
		Password:    GetFlagValue(cmd, "password"),
		Description: GetFlagValue(cmd, "description"),
	}

	err := db.DBConnection.Create(&user)
	if err != nil {
		log.Printf("error in adding the password: %v\n", err.Error())
		return
	}

	log.Printf("password saved successfully: %v\n", user.ID)
}

func EditPassword(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	description, _ := cmd.Flags().GetString("description")

	// Use passed-in data here
	fmt.Printf("Adding new password - Title: %s, Username: %s, Password: %s, Description: %s\n", title, username, password, description)
}

func DeletePassword(cmd *cobra.Command, args []string) {
	id := GetFlagValue(cmd, "id")
	err := db.DBConnection.Delete(id)
	if err != nil {
		fmt.Printf("error in delete the user: %v\n", err.Error())
		return
	}
	fmt.Printf("User removed successfully: %v\n", id)
}

func GetPassword(cmd *cobra.Command, args []string) {
	id := GetFlagValue(cmd, "id")

	user, err := db.DBConnection.Get(id)
	if err != nil {
		fmt.Printf("error in getting the user: %v\n", err.Error())
		return
	}

	// Use passed-in data here
	fmt.Printf("Adding new password - Title: %s, Username: %s, Password: %s, Description: %s\n", user.Title, user.Username, user.Password, user.Description)
}

func ListPassword(cmd *cobra.Command, args []string) {

	users, err := db.DBConnection.List()
	if err != nil {
		fmt.Printf("error in getting the users: %v\n", err.Error())
		return
	}

	for _, user := range users {
		fmt.Printf("Adding new password - Title: %s, Username: %s, Password: %s, Description: %s\n", user.Title, user.Username, user.Password, user.Description)
	}
}
