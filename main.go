package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	const logFileName = "time_entries.log"

	// Open the log file for writing (overwrite any previous content)
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	fmt.Println("Enter time in 'hours:minutes' format (e.g., 1:30). Press Ctrl+C to finish and display the total.")
	fmt.Printf("Previous entries will be overwritten in '%s'.\n", logFileName)

	totalMinutes := 0
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for scanner.Scan() {
			input := scanner.Text()
			parts := strings.Split(input, ":")
			if len(parts) != 2 {
				fmt.Println("Invalid format. Please enter time in 'hours:minutes' format.")
				continue
			}

			hours, err1 := strconv.Atoi(parts[0])
			minutes, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid numbers. Please try again.")
				continue
			}

			_, writeErr := logFile.WriteString(input + "\n")
			if writeErr != nil {
				fmt.Println("Failed to write to log file:", writeErr)
				continue
			}

			totalMinutes += hours*60 + minutes
		}
	}()

	<-sigChan
	fmt.Printf("\nCalculation finished. The total time is: ")
	fmt.Printf("%d:%02d\n", totalMinutes/60, totalMinutes%60)
	fmt.Printf("The final input was saved to '%s'.\n", logFileName)
}
