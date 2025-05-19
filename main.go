package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter a text to analyze: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if len(text) == 0 {
			fmt.Println("\nText cannot be empty. Please try again.")
			continue
		}

		wordCount := len(strings.Fields(text))
		charCount := len(strings.ReplaceAll(text, " ", ""))

		var letter string
		for {
			fmt.Print("Enter a letter to count: ")
			fmt.Scanln(&letter)
			if len(letter) == 1 && ((letter[0] >= 'A' && letter[0] <= 'Z') || (letter[0] >= 'a' && letter[0] <= 'z')) {
				break
			}
			fmt.Println("\nPlease enter a single letter (a-z or A-Z).")
		}

		count := countLetterOccurrences(text, letter)

		fmt.Println("\n--- Analysis Result ---")
		fmt.Println("Characters (excluding spaces):", charCount)
		fmt.Println("Words:", wordCount)
		fmt.Printf("Occurrences of '%s': %d\n", letter, count)

		fmt.Print("\nDo you want to analyze another text? (yes/no): ")
		again, _ := reader.ReadString('\n')
		again = strings.ToLower(strings.TrimSpace(again))
		if again != "yes" {
			fmt.Println("Goodbye!")
			break
		}
	}
}

func countLetterOccurrences(text, letter string) int {
	count := 0
	lowerText := strings.ToLower(text)
	lowerLetter := strings.ToLower(letter)

	for _, ch := range lowerText {
		if string(ch) == lowerLetter {
			count++
		}
	}

	return count
}
