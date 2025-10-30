package main

import (
	"fmt"
)

func main() {
	servers := []string{"AWS", "Azure"}
	fmt.Println("Initial Servers: ", servers)
	fmt.Println("Length: ", len(servers), " | Capacity: ", cap(servers))

	// Append more servers
	servers = append(servers, "DigitalOcean", "Oracle", "Citrix")
	fmt.Println("After Scaling Up (append): ", servers)
	fmt.Println("Length: ", len(servers), " | Capacity: ", cap(servers))

	// Accessing elements
	fmt.Println("Accessing Examples: ")
	fmt.Println("First Server from Slice: ", servers[0])
	fmt.Println("Last Server from Slice: ", servers[len(servers)-1])

	// Sub-slice
	subSlice := servers[1:3]
	fmt.Println("Sub Slice (Mid Environments): ", subSlice)

	// Remove one server
	removedIndex := 2
	servers = append(servers[:removedIndex], servers[removedIndex+1:]...)
	fmt.Println("After removing one server (Scale Down): ", servers)
}
