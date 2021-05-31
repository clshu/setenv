package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	var name string
	var args []string

	if len(os.Args) < 3 {
		fmt.Println("setenv dotenvfile [go|fresh]")
		os.Exit(1)
	}
	dotenvfile := os.Args[1]

	envMap, err := godotenv.Read(dotenvfile)
	if err != nil {
		fmt.Printf("Reading file %v failed. %v", dotenvfile, err.Error())
		os.Exit(2)
	}

	if envMap["NODE_ENV"] != "development" && envMap["NODE_ENV"] != "test" {
		fmt.Println("It is not development or test environment")
		os.Exit(0)
	}
	switch os.Args[2] {
	case "go":
		name = "go run"
		args = []string{"server.go"}
		break
	case "fresh":
		name = "fresh"
		args = []string{}
		break
	default:
		fmt.Println("setenv dotenvfile [go|fresh]")
		os.Exit(3)
	}

	cmd := exec.Command(name, args...)

	for key, value := range envMap {
		str := key + "=" + value
		cmd.Env = append(os.Environ(), str)
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
