package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file name Argument")
		os.Exit(1)
	}
	fileName := os.Args[1]

	envMap, err := godotenv.Read(fileName)
	if err != nil {
		fmt.Printf("Reading file %v failed. %v", fileName, err.Error())
		os.Exit(2)
	}	
	
	for key, value := range envMap {
		str := "export " + key + "=" + value
		cmd := exec.Command(str)
		cmd.Run()
	}	
}