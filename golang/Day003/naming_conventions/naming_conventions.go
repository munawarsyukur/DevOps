package main

import "fmt"

func main() {
	// Exported variable
	var ExportedVariable int = 42
	fmt.Println("Exported Variable:", ExportedVariable)

	// Non-exported variable
	var nonExportedVariable int = 99
	fmt.Println("Non-exported Variable:", nonExportedVariable)
}
