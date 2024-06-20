// Assignment 1
// learned all topics from go offical site and apply it into below code of hangman game.
// Understand variable declaration standard output print functions
// Learn basic condition statements and data types
// Understand basic looping structure
// Initialize entries map
// Show placeholder
// Basic conditions/evaluations
// Print guessed letters

package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"

  // lookup for entries made by the user.
	entries :=  map[string]bool{} 

  // list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}

	for i:= 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	chances := len(word)

	for {
		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		userInput := strings.Join(placeholder,"")
		if chances == 0 && userInput != word {
			fmt.Println("Game over! Try agin")
			break
		}

		// evaluate a win!
		if userInput == word {
			fmt.Println("You Win!!")
			break
		}

    // Console display
		fmt.Println(placeholder) // render the placeholder
		fmt.Printf("Chances:%d\n",chances) // render the chances left

		keys := []string{}
		for k, _ := range entries {
			keys = append(keys, k)
		}

		fmt.Printf("words guessed till now: %v \n", keys) // show the letters or words guessed till now.
		fmt.Printf("Guess a letter or the word: ")

    // take the input
		str := ""
		fmt.Scanln(&str)
		fmt.Printf("Guessed letter: %v \n", str)

    // compare and update entries, placeholder and chances.
  }
}