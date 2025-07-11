package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Notes struct {
	ID        int
	Content   string
	CreatedAt time.Time
}

var AllNotes []Notes

// add note
func AddNote(notes Notes) {
	AllNotes = append(AllNotes, notes)
	fmt.Println("Notes Added Successfully")
}

// delete <id> – delete a note by its ID
func Delete(id int) {
	if id < 0 || id >= len(AllNotes) {
		fmt.Println("Add a Correct ID, you wanna delete")
		return
	}
	AllNotes = append(AllNotes[:id], AllNotes[id+1:]...)
	fmt.Println("Note Deleted")
}

// List Nodes
func ListNotes() {
	for _, read := range AllNotes {
		fmt.Printf("ContentID: %d , Content: %s\n", read.ID, read.Content)
	}
	fmt.Println("All notes printed")
}

// search
func SearchNotes(keyword string) {
	found := false
	for _, read := range AllNotes {
		if strings.Contains(strings.ToLower(read.Content), strings.ToLower(keyword)) {
			fmt.Printf("Note Found: %d , %s, %s", read.ID, read.Content, read.CreatedAt)
			found = true
		}
	}
	if !found {
		fmt.Println("No matching notes found.")
	}
}

// Save
func SaveData() {
	b, err := json.MarshalIndent(AllNotes, "", "")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("notes.json", b, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Notes saved to notes.json.")
}

func LoadData() {
	b, err := os.ReadFile("notes.json")
	if err != nil {
		if os.IsNotExist(err) {
			AllNotes = []Notes{}
			return
		}
		panic(err)
	}
	if len(b) == 0 {
		fmt.Println("notes.json is empty. Starting with no notes.")
		AllNotes = []Notes{}
		return

	}
	err = json.Unmarshal(b, &AllNotes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Notes loaded from notes.json.")
}
