package main

import (
	"fmt"
	"os"
)

// DEFER Keyword is to use to do something right after the function executed or at the end of the function
// for example closing the file after opening

// one use case

// This function shows the real use of 'defer' for cleanup.
func createFileAndDeferClose(fileName string) {
	file, err := os.Create(fileName + ".txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	//IMPORTANT: We use 'defer' to schedule the file closing.
	// This command will run when createFileAndDeferClose() finishes,
	// no matter if it finishes normally or with an error.
	defer file.Close()

	_, err = file.WriteString("User data processed successfully!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	// Because of 'defer', we do not need to call file.Close() here.
	fmt.Println("\n-- File Operation --")
	fmt.Println("File " + fileName + " was opened, written to, and is scheduled to be closed by 'defer'.")
}

func main() {

	createFileAndDeferClose("Hello-Differ")
}
