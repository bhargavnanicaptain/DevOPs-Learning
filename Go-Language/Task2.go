package main

import (
	"fmt"
	"strings"
)

// 1️⃣ Function to analyze a single log line
func analyzeLine(line string) string {
	if strings.Contains(line, "ERROR") {
		return "Error"
	} else if strings.Contains(line, "SUCCESS") {
		return "Success"
	} else {
		return "Info"
	}
}

// 2️⃣ Function to process all log entries and return results
func processLogs(logs []string) (int, int, int) {
	errorCount := 0
	successCount := 0
	infoCount := 0

	for _, line := range logs {
		category := analyzeLine(line)
		switch category {
		case "Error":
			fmt.Println("❌", line)
			errorCount++
		case "Success":
			fmt.Println("✅", line)
			successCount++
		case "Info":
			fmt.Println("ℹ️", line)
			infoCount++
		}
	}

	return errorCount, successCount, infoCount
}

// 3️⃣ Main function — orchestrates everything
func main() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("🧩 SONATA SOFTWARE — Deployment Log Analyzer")
	fmt.Println("---------------------------------------------------")

	logs := []string{
		"INFO: Deployment started",
		"INFO: Connecting to database",
		"ERROR: Connection timeout",
		"INFO: Retrying...",
		"SUCCESS: Deployment completed",
		"INFO: Cleaning up temp files",
		"ERROR: Unable to delete cache",
	}

	errors, success, info := processLogs(logs)

	fmt.Println("---------------------------------------------------")
	fmt.Println("📊 Log Summary Report")
	fmt.Println("ℹ️ Info Messages:", info)
	fmt.Println("❌ Errors:", errors)
	fmt.Println("✅ Success:", success)

	if errors > 0 {
		fmt.Println("⚠️  Deployment completed with errors! Immediate review needed.")
	} else {
		fmt.Println("🚀 Deployment successful! No errors found.")
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("✅ Log analysis completed successfully by Bhargav ❤️")
}
