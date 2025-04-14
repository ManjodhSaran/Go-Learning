package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo/todo"
)

type saver interface {
	Save() error
}
type displayer interface {
	saver
	Display()
}

func main() {
	text := getUserInput("Enter text: ")
	_todo, err := todo.New(text)
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}
	saveData(_todo)
	displayData(_todo)
	fmt.Println("Todo saved successfully!")
}

func saveData(data saver) {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
	}
}

func displayData(data displayer) {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
	}
	data.Display()
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
