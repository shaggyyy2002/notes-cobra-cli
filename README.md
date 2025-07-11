## 📒 Notes CLI

A simple command-line notes app in Go where you can:

1. Add notes
2. List notes with timestamps and IDs
3. Delete notes by ID
4. Search notes by keyword
5. Persist notes to a JSON file

---

## 🛠 Features

### 1️⃣ Add a Note

Add a new note with:

```
notes add "Your note content here"
```

* Automatically assigns a unique ID.
* Adds a timestamp of when the note was created.

---

### 2️⃣ List Notes

View all saved notes:

```
notes list
```

* Displays:

  * ID
  * Note content
  * Created timestamp

---

### 3️⃣ Delete a Note

Delete a note by its ID:

```
notes delete 2
```

* Removes the note with the given ID from your notes file.

---

### 4️⃣ Search Notes

Search for notes containing a keyword:

```
notes search "meeting"
```

* Displays all notes that contain the keyword (case-insensitive).

---

## 🗂 How Data is Stored

* Notes are stored in a `notes.json` file in the same directory.
* Each note has:

  ```go
  type Note struct {
      ID        int
      Content   string
      CreatedAt time.Time
  }
  ```
* The CLI reads this file at startup and writes to it when notes are added or deleted.

## Available Commands:
  - `AddNote`     This will allow you to add note
  - `Delete`      This will delete your note
  - `ListNotes`   A brief description of your command
  - `LoadData`    Data is Loaded
  - `SaveData`    Saves Data in the JSON file
  - `SearchNotes` A brief description of your command
  - `completion`  Generate the autocompletion script for the specified shell
  - `help`       Help about any command

---