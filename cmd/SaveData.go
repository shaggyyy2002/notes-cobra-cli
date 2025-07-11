/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/shaggyyy2002/notes-cli/note"
	"github.com/spf13/cobra"
)

// SaveDataCmd represents the SaveData command
var SaveDataCmd = &cobra.Command{
	Use:   "SaveData",
	Short: "Saves Data in the JSON file",
	Long:  `This will save your current data to the JSON File`,
	Run: func(cmd *cobra.Command, args []string) {
		note.SaveData()
		fmt.Println("Data is Saved")
	},
}

func init() {
	rootCmd.AddCommand(SaveDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SaveDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SaveDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
