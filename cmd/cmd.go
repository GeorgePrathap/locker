package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GenerateCmd() {
	rootCmd := &cobra.Command{
		Use:   "locker",
		Short: "Cmd based password locker",
	}

	rootCmd.AddCommand(getAddCmd(), getEditCmd(), getListCmd(), getCmd(), deleteCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func deleteCmd() *cobra.Command {
	delete := &cobra.Command{
		Use:   "delete",
		Short: "delete the password",
		Run:   DeletePassword,
	}

	delete.Flags().StringP("id", "i", "", "ID for the password (required)")
	delete.MarkFlagRequired("id")
	return delete
}

func getCmd() *cobra.Command {
	get := &cobra.Command{
		Use:   "get",
		Short: "Get all password",
		Run:   GetPassword,
	}

	get.Flags().StringP("id", "i", "", "ID for the password (required)")
	get.MarkFlagRequired("id")
	return get
}

func getListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all password",
		Run:   ListPassword,
	}
}

func getEditCmd() *cobra.Command {
	edit := &cobra.Command{
		Use:   "edit",
		Short: "Edit the existing password",
		Run:   EditPassword,
	}

	edit.Flags().StringP("title", "t", "", "Title of the new password (required)")
	edit.Flags().StringP("username", "u", "", "Username of the new password (required)")
	edit.Flags().StringP("password", "p", "", "Password of the new password (required)")
	edit.Flags().StringP("description", "d", "", "Description of the new password (optional)")
	edit.Flags().StringP("id", "i", "", "ID for the password (required)")

	edit.MarkFlagRequired("id")
	return edit
}

func getAddCmd() *cobra.Command {
	add := &cobra.Command{
		Use:   "add",
		Short: "Add the new password",
		Run:   AddPassword,
	}

	add.Flags().StringP("title", "t", "", "Title of the new password (required)")
	add.MarkFlagRequired("title")

	add.Flags().StringP("username", "u", "", "Username of the new password (required)")
	add.MarkFlagRequired("username")

	add.Flags().StringP("password", "p", "", "Password of the new password (required)")
	add.MarkFlagRequired("password")

	add.Flags().StringP("description", "d", "", "Description of the new password (optional)")
	return add
}

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("title", "t", "", "Title of the new password (required)")
	cmd.MarkFlagRequired("title")

	cmd.Flags().StringP("username", "u", "", "Username of the new password (required)")
	cmd.MarkFlagRequired("username")

	cmd.Flags().StringP("password", "p", "", "Password of the new password (required)")
	cmd.MarkFlagRequired("password")

	cmd.Flags().StringP("description", "d", "", "Description of the new password (optional)")
}
