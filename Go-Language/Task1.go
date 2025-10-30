package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Constants
	const Company = "Sonata Software"
	const SupportedEnvs = "Dev | QA | Prod"

	// 2️⃣ Variables
	appName := "DigitalWellbeingTracker"
	version := 1.2
	build := 1004
	deployedBy := "Bhargav"
	environment := "QA"
	deploymentSuccess := true

	// 3️⃣ Print deployment summary
	fmt.Println("---------------------------------------------------")
	fmt.Println("🏢 Company:", Company)
	fmt.Println("📦 Application:", appName)
	fmt.Printf("🔖 Version: %.1f (Build #%d)\n", version, build)
	fmt.Println("👨‍💻 Deployed By:", deployedBy)
	fmt.Println("🌍 Environment:", environment)
	fmt.Println("💡 Supported Environments:", SupportedEnvs)
	fmt.Println("---------------------------------------------------")

	// 4️⃣ Simulate deployment status
	if deploymentSuccess {
		fmt.Println("✅ Deployment Successful!")
	} else {
		fmt.Println("❌ Deployment Failed!")
	}

	// 5️⃣ Update variables to simulate next release
	fmt.Println("\n🚀 Preparing next release...")
	version += 0.1
	build += 1
	environment = "Production"
	deploymentSuccess = false // simulate a failure

	// 6️⃣ Print updated release info
	fmt.Println("---------------------------------------------------")
	fmt.Printf("🔖 New Version: %.1f (Build #%d)\n", version, build)
	fmt.Println("🌍 Target Environment:", environment)
	fmt.Println("👨‍💻 Released By:", deployedBy)
	fmt.Println("---------------------------------------------------")

	if !deploymentSuccess {
		fmt.Println("⚠️  Production deployment failed, initiating rollback...")
		// Rollback simulation
		version -= 0.1
		build -= 1
		environment = "QA"
		fmt.Printf("🔁 Rolled back to version %.1f (Build #%d) in %s\n", version, build, environment)
	}

	fmt.Println("✅ Version tracking completed successfully!")
}
