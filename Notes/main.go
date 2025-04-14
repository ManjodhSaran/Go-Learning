package main

import (
	"bufio"
	"errors"
	"fmt"
	"notes/note"
	"os"
	"strings"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		panic(err)
	}

	_note, err := note.New(title, content)
	if err != nil {
		panic(err)
	}

	_note.Display()
	err = _note.Save()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved")
}

func getNoteData() (string, string, error) {
	title := getUserInput("Enter title: ")
	content := getUserInput("Enter content: ")

	if title == "" {
		return "", "", errors.New("title is required")
	}
	if content == "" {
		return "", "", errors.New("content is required")
	}

	return title, content, nil
}

func getUserInput(str string) string {
	fmt.Print(str)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
