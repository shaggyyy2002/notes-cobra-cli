/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"
	"time"

	"github.com/shaggyyy2002/notes-cli/note"
	"github.com/spf13/cobra"
)

// AddNoteCmd represents the AddNote command
var AddNoteCmd = &cobra.Command{
	Use:   "AddNote",
	Short: "This will allow you to add note",
	Long: `AddNote allows you to crate a note for yourself. It Automatically assigns a unique ID and 
	Adds a timestamp of when the note was created.`,
	Run: func(cmd *cobra.Command, args []string) {
		content := strings.Join(args, " ")

		newNote := note.Notes{
			ID:        len(note.AllNotes),
			CreatedAt: time.Now(),
			Content:   content,
		}
		note.AddNote(newNote)
		note.SaveData()
	},
}

func init() {
	rootCmd.AddCommand(AddNoteCmd)
}
