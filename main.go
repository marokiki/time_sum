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

	fmt.Println("Enter time in 'hours:minutes' format (e.g., 1:30). Press Ctrl+C to finish and display the total.")

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

			totalMinutes += hours*60 + minutes
		}
	}()

	<-sigChan
	fmt.Println("\nCalculation finished. The total time is:")
	fmt.Printf("%d:%02d\n", totalMinutes/60, totalMinutes%60)
}
