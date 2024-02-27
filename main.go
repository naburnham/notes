package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func loadConfig() map[string]string {

	config := make(map[string]string)

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}
	home_dir, err := os.UserHomeDir()

	directory := os.Getenv("DIRECTORY")
	notebook := os.Getenv("NOTEBOOK")

	config["home_dir"] = home_dir
	config["directory"] = directory
	config["notebook"] = notebook

	return config
}

func setup(route, notebook string) {
	if _, err := os.Stat(route); os.IsNotExist(err) {
		err := os.Mkdir(route, 0766)
		if err != nil {
			log.Printf("Failed to make directory: %v - %v\n", route, err)
		}
	}

	if _, err := os.Stat(notebook); os.IsNotExist(err) {
		f, err := os.Create(notebook)
		if err != nil {
			log.Printf("Failed to create notebook: %v", err)
		}
		f.Close()
	}
}

func getNotebooks(route string) {
	files, err := os.ReadDir(route)
	if err != nil {
		log.Printf("Failed to find notebooks. %v", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func setNotebook(route, filename string) {

}

func createBook(route, filename string) {
	filepath := filepath.Join(route, filename)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		f, err := os.Create(filepath)
		if err != nil {
			log.Printf("Failed to create notebook: %v", err)
		}
		f.Close()
	}
}

func addNote(file string) {

}

func main() {
	configs := loadConfig()
	route := filepath.Join(configs["home_dir"], configs["directory"])
	notebook := filepath.Join(route, configs["notebook"])
	setup(route, notebook)

	if len(os.Args) == 1 {
		os.Exit(0)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("help")
		break
	case "showbooks":
		getNotebooks(route)
		break
	case "setbook":
		setNotebook(route, os.Args[2])
		break
	case "createbook":
		if len(os.Args) < 3 {
			fmt.Println("Need a name for the book")
			os.Exit(1)
		}
		createBook(route, os.Args[2])

	default:
		fmt.Println("Invalid command")
	}
}
