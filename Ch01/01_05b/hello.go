package main

import(
	"fmt"
)

//Creating a go file
//required a package, import and and func main, as this file is setup
//creating a module its required
//terminal step go mod init com.example/hello -> creates a go.mod file
//Creating a executable its required
//Executable go build -o hello hello.go -> creates a hello file that can be run in terminal with ./hello
func main(){
	fmt.Println("Test to Print Hello world")
}