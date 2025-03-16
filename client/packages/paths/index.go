package paths

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

func Bootstrap() {
	err := godotenv.Load()

	if err != nil {
		panic("could not load .env file")
	}

	if !CreateConfigDirectory() {
		panic("config directory does not exists and could not be made")
	}

	if !CreatePathsFile() {
		panic("config file does not exists and could not be made")
	}
}

func Manage() {
	var action int
	for {
		if action == 6 {
			break
		}

		color.Green("Please choose what you want to do:")
		color.Yellow("1) View current path")
		color.Yellow("2) Set path")
		color.Yellow("6) Exit application")

		_, err := fmt.Scan(&action)

		if err != nil {
			CouldNotReadAction()
		}

		screen.Clear()

		switch action {
		case 1:
			List()
		case 2:
			Set()
		case 6:
			Exit()
		default:
			CouldNotReadAction()
		}
	}
}

// List all the current registered paths.
func List() {
	contents, err := GetContents()

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(contents.Paths) == 0 {
		color.Red("No paths registered yet")
		return
	}

	fmt.Println("Current registered paths for sync are:")

	for index, path := range contents.Paths {
		fmt.Println(strings.Join([]string{
			strconv.Itoa(index), path}, " "))
	}
}

// Set path to the config file.
func Set() {
	var pathInput string

	color.Green("Please input your path you'd like to sync")
	fmt.Scan(&pathInput)

	_, err := os.ReadDir(pathInput)

	if err != nil {
		color.Red("Could not read path. Check path or input another one")
		fmt.Println(err)
		return
	}

	contents, err := GetContents()

	if err != nil {
		color.Red("Could not read contents of pathfile. Please check pathfile")
		return
	}

	contents.Paths = []string{pathInput}
	err = WriteContents(contents)

	if err != nil {
		color.Red("Could not write to path file")
		return
	}

	color.Green("Succesfully written to paths")
	return
}

// Exit application
func Exit() {
	fmt.Println("Bye")
}

// CouldNotReadAction Panic response if the action could not be read.
func CouldNotReadAction() {
	fmt.Println("Could not read action. Please select out of 1, 2 or 6")
}
