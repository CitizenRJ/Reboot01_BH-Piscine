package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main(){
	arguments := os.Args[1:]
	if len(os.Args[1:]) == 0{
		fmt.Println("File name missing")
	} else if len (os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
	} else {
		fileData, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(fileData))
	}
}
