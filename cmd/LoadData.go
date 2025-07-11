/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/shaggyyy2002/notes-cli/note"
	"github.com/spf13/cobra"
)

// LoadDataCmd represents the LoadData command
var LoadDataCmd = &cobra.Command{
	Use:   "LoadData",
	Short: "Data is Loaded",
	Long:  `All the data is loaded to the Slice and can be used again`,
	Run: func(cmd *cobra.Command, args []string) {
		note.LoadData()
		fmt.Println("Data is Loaded, you can see all in the list")
	},
}

func init() {
	rootCmd.AddCommand(LoadDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// LoadDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// LoadDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
