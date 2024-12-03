package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note_app/note"
)

func main() {

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the node failed.")
		return
	}

	fmt.Println("Saving the note succeeded")
}

func getNoteData() (string, string) {
	// get user input
	noteTitle := getUserInput("Enter Note Title: ")

	noteContent := getUserInput("Enter Note Content: ")

	return noteTitle, noteContent

}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
