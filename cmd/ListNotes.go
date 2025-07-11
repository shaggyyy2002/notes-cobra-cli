/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/shaggyyy2002/notes-cli/note"
	"github.com/spf13/cobra"
)

// ListNotesCmd represents the ListNotes command
var ListNotesCmd = &cobra.Command{
	Use:   "ListNotes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		note.ListNotes()
	},
}

func init() {
	rootCmd.AddCommand(ListNotesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListNotesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListNotesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
