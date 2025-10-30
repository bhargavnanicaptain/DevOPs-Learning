package main

import "fmt"

func greetUser(name string) {
	fmt.Println("Helo", name, "Welcome to Devops World!")
}

func add(a int, b int) int {
	return a + b
}

func serverStatus(server string) (string, bool) {
	return server, true
}
func main() {
	fmt.Println("---------------------")
	greetUser("Bhargav")
	fmt.Println("---------------------")
	sum := add(5, 10)
	fmt.Println("Sum: ", sum)
	fmt.Println("---------------------")
	name, status := serverStatus("AWS")
	fmt.Println("Server: ", name, "| Active: ", status)
}
