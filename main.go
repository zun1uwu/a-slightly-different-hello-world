package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func InitFile(fileName, fileContent, finalLocation string) {

	// creates the file and sends a log to the console
	File, errCreate := os.Create(fileName)
	if errCreate != nil {
		log.Fatal(errCreate)
	}

	log.Println(File)

	// writes content to the file, checks if successful and closes it afterwards
	_, errWS := File.WriteString(fileContent)
	if errWS != nil {
		log.Fatal(errWS)
	}

	File.Close()

	// gets the path of where the executable is ran from
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	CurrentPath := filepath.Dir(ex)

	// uses the current path of the created file,
	// which is next to the executable by default,
	// then moves the file to it's final location
	errMove := os.Rename(CurrentPath+"\\"+fileName, finalLocation)
	if errMove != nil {
		fmt.Println(errMove)
	}

	defer File.Close()

}

func main() {

	// places of interest in a windows filesystem
	HomePath := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	fmt.Println(HomePath)
	AppdataPath := HomePath + "\\AppData"
	fmt.Println(AppdataPath)
	DesktopPath := HomePath + "\\Desktop"

	// creates a file using the function "InitFile" above for easier handling
	InitFile("emptyFile.txt", "Hello world!", DesktopPath+"\\File.txt")

}
