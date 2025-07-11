/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/shaggyyy2002/notes-cli/note"
	"github.com/spf13/cobra"
)

// DeleteCmd represents the Delete command
var DeleteCmd = &cobra.Command{
	Use:   "Delete",
	Short: "This will delete your note",
	Long:  `Using this command this will delete your note from the JSON File using the ID`,
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid ID")
		}
		note.Delete(id)
		note.SaveData()
	},
}

func init() {
	rootCmd.AddCommand(DeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
