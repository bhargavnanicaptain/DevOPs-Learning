package main

import (
	"fmt"
)

func main() {
	var servers [3]string = [3]string{"AWS", "Azure", "GCP"}
	fmt.Println("Cloud Servers: ", servers)
	//Printing Using Loops
	for index, server := range servers {
		fmt.Println("Server ", index+1, server)
	}
}
