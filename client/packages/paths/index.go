package paths

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"os"
	"strconv"
	"strings"
)

func Bootstrap() {
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
		color.Yellow("1) View current paths")
		color.Yellow("2) Add to paths")
		color.Yellow("3) Remove from paths")
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
			Add()
		case 3:
			Remove()
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

// Add input to the path file.
func Add() {
	var pathInput string

	color.Green("Please input your path you'd like to sync")
	fmt.Scan(&pathInput)

	_, err := os.ReadDir(pathInput)

	if err != nil {
		color.Red("Could not read path. Check path or input another one")
		return
	}

	contents, err := GetContents()

	if err != nil {
		color.Red("Could not read contents of pathfile. Please check pathfile")
		return
	}

	contents.Paths = append(contents.Paths, pathInput)
	err = WriteContents(contents)

	if err != nil {
		color.Red("Could not write to path file")
		return
	}

	color.Green("Succesfully written to paths")
	return
}

// Remove PATH from list
func Remove() {
	var inputIndex int

	color.Green("Please input index number of path you'd like to remove.")
	color.White("To get index, view all paths first.")

	_, err := fmt.Scan(&inputIndex)

	if err != nil {
		color.Red("Please enter a number. Returning to menu")
		return
	}

	contents, err := GetContents()

	if err != nil {
		color.Red("Could not read contents of pathfile. Please check pathfile")
		return
	}

	if inputIndex < 0 || inputIndex > len(contents.Paths) {
		color.Red("Given index is not within the range of pathlist. Please view pathfile first")
		return
	}

	contents.Paths = append(contents.Paths[:inputIndex], contents.Paths[inputIndex+1:]...)
	err = WriteContents(contents)

	if err != nil {
		color.Red("Could not write to path file")
		return
	}

	color.Green("Succesfully removed from paths")
}

// Exit application
func Exit() {
	fmt.Println("Bye")
}

// CouldNotReadAction Panic response if the action could not be read.
func CouldNotReadAction() {
	fmt.Println("Could not read action. Please select out of 1, 2, 3 or 6")
}
