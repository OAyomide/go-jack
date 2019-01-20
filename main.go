package main

import (
	"flag"
	"fmt"
	"os"

	logger "github.com/go-jack/logger"
)

//Bot struct is the struct our bot property will be made after
//BRAINSTORM: SHouldnt it be an interface? 🤔
type Bot struct {
	name    string
	channel string
}

// func init() {
// 	logger.CreateLogFolder("error")
// 	logger.CreateLogFile("error", "error.log")
// }

// declare our CLI flags
var (
	mainCommand = flag.NewFlagSet("new", flag.ExitOnError)

	// // create subcommand for command new
	name = mainCommand.String("name", "", "Creates a new project with the specified name")
	full = mainCommand.Bool("full", true, "Create a project with all config out of the box")
)

func main() {

	// make sure the subcommand has been passed
	if len(os.Args) < 2 {
		fmt.Println("Please specify a project structure to create")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		mainCommand.Parse(os.Args[2:])

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check which command was parsed
	if mainCommand.Parsed() {
		if *name == "" {
			fmt.Println("Specify a name for the new project")
			mainCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	createScaffold(*name)
}

// this function creates the sample bot based on the https://github.com/paked/messenger/example/basic/main.go
func createScaffold(foldername string) {
	// we want to create the folder first

	err := os.Mkdir(foldername, 0777)

	if err != nil {
		logger.Print("error", "Couldn't create folder for the new project", err.Error())
	}

	// we want to copy the files from our example
	file, err := os.Create(foldername + "/" + "main.go")
	if err != nil {
		logger.Print("error", "Error creating the main.go file", err.Error())
	}

	defer file.Close()

	// we want to read the main.go file we're copying from
	fl, err := os.Open("example/main.go")
	if err != nil {
		logger.Print("error", "ERROR OPENING THE EXAMPLE main.go FILE", err.Error())
		return
	}
	defer fl.Close()

	flInfo, err := fl.Stat()
	if err != nil {
		logger.Print("error", "ERROR GETTING STATS FOR THE OPENED FILE", err.Error())
		return
	}

	flSize := flInfo.Size()
	fileBuffer := make([]byte, flSize)

	bytesRead, err := fl.Read(fileBuffer)
	if err != nil {
		logger.Print("error", "ERROR READING FILE", err.Error())
		return
	}

	fmt.Printf("READ THE FOLLOWING FROM THE FILE\n", string(bytesRead))

	// we want to write the file copied from example/main.go to testbot/main.go
	_, er := file.Write(fileBuffer)

	if er != nil {
		logger.Print("ERROR WRITING TO THE FILE.", err.Error())
	}
}
