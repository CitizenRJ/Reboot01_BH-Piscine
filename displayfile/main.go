package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(os.Args[1:]) == 0 {
		fmt.Print("File name missing")
		fmt.Print("\n")
	} else if len(os.Args[1:]) > 1 {
		fmt.Print("Too many arguments")
		fmt.Print("\n")
	} else {
		fileData, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(string(fileData))
	}
}
