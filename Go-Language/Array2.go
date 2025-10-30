package main

import (
	"fmt"
)

func main() {
	var servers [4]string
	fmt.Println("Enter the 4 server names")
	for i := 0; i < len(servers); i++ {
		fmt.Println("Servers: ", i+1)
		fmt.Scan(&servers[0], &servers[1], &servers[2], &servers[3])

	}

	fmt.Println("Servers List: ")
	for i := 0; i < len(servers); i++ {
		fmt.Println("Servers ", i+1, servers[i])
	}
}
