package main

import (
	"fmt"
	"strings"
)

func main() {
	printWelcomeBanner()

	for {
		printMenu()
		choice := getUserChoice()

		if choice == 0 {
			printGoodbye()
			break
		}

		executeChoice(choice)

		// Pause before showing menu again
		fmt.Println("\n" + strings.Repeat("─", 60))
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}

func printWelcomeBanner() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                        ║")
	fmt.Println("║          🚀  WELCOME TO GO TUTORIAL  🚀               ║")
	fmt.Println("║                                                        ║")
	fmt.Println("║         Learn Go Programming Step by Step             ║")
	fmt.Println("║                                                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Println(strings.Repeat("═", 60) + "\n")
}

func printMenu() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("📚  AVAILABLE TOPICS")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println()

	topics := []string{
		"Variables",
		"Constants",
		"Data Types",
		"Arrays",
		"Slices",
		"Operators",
		"Conditions",
		"Loops",
		"Functions",
		"Structs",
		"Maps",
		"Defer",
	}

	for i, topic := range topics {
		fmt.Printf("  %2d. %-20s", i+1, topic)
		if (i+1)%2 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println(strings.Repeat("─", 60))
	fmt.Println("  0. Exit Tutorial")
	fmt.Println(strings.Repeat("═", 60))
	fmt.Print("\n👉 Enter your choice: ")
}

func getUserChoice() int {
	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		var discard string
		fmt.Scanln(&discard)
		return -1
	}

	return choice
}

func executeChoice(choice int) {
	fmt.Println()

	switch choice {
	case 1:
		variables()
	case 2:
		constants()
	case 3:
		dataTypes()
	case 4:
		arrays()
	case 5:
		slices()
	case 6:
		operators()
	case 7:
		conditions()
	case 8:
		loops()
	case 9:
		functions()
	case 10:
		structs()
	case 11:
		maps()
	case 12:
		defers()
	default:
		fmt.Println("❌ Invalid choice! Please enter a number between 0 and 12.")
	}
}

func printGoodbye() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                        ║")
	fmt.Println("║           ✨  Thank You for Learning Go!  ✨          ║")
	fmt.Println("║                                                        ║")
	fmt.Println("║              Keep Coding and Have Fun! 🎉             ║")
	fmt.Println("║                                                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Println(strings.Repeat("═", 60) + "\n")
}
