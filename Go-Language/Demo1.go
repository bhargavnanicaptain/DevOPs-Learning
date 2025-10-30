package main

import (
	"fmt"
)

func main() {
	//Decalring Variables as Constants
	const companyName = "Sonata Software"
	const hoursPerDay = 24

	//Declaring Variables with Explicit Type
	var serverName string = "aws-prod-server"
	var region string = "ap-south-1"
	var costPerHour float64 = 0.45
	var isActive bool = true
	var upTimeDays int = 120

	//Declaring Variables using Short hand :=
	owner := "Bhargav"
	environment := "Production"

	//Perform some caluclations
	totalHours := upTimeDays * hoursPerDay
	totalCost := float64(totalHours) * costPerHour

	//Type Conversion Example
	upTimeInWeeks := float64(upTimeDays) / 7.0

	//Printing the Values
	fmt.Println("-------------------------------------------------")
	fmt.Println("🏢 Company:", companyName)
	fmt.Println("👨‍💻 Server Owner:", owner)
	fmt.Println("💻 Server Name:", serverName)
	fmt.Println("🌍 Region:", region)
	fmt.Println("⚙️ Environment:", environment)
	fmt.Println("🔌 Active:", isActive)
	fmt.Printf("📅 Uptime: %d days (%.2f weeks)\n", upTimeDays, upTimeInWeeks)
	fmt.Printf("💰 Cost per Hour: $%.2f\n", costPerHour)
	fmt.Printf("💵 Total Cost till now: $%.2f\n", totalCost)
	fmt.Println("-------------------------------------------------")
}
