package main // package can have a lot of files inside of it. a bundle. a module.
// two types: executable or helper

import "fmt" // give my package main access to the fmt package

// must have a function called main - this defines an executable package
func main() {
	fmt.Println("Hello, World!")
}
