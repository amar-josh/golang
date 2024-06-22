// Assignment 2

package main

import (
	"fmt"
	"strings"
)

func get_keys(entries map[string]bool)(keys []string){
		for k, _ := range entries {
			keys = append(keys, k)
		}
	return
}

func main() {
	word := "elephant"

  // lookup for entries made by the user.
	entries :=  map[string]bool{} 

  // list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	// placeholder := make([]string, len(word), len(word))

	for i:= 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	chances := 8

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
		fmt.Println();
		fmt.Println(placeholder) // render the placeholder
		fmt.Printf("Chances:%d\n",chances) // render the chances left
		fmt.Printf("words guessed till now: %v \n", get_keys(entries)) // show the letters or words guessed till now.
		fmt.Printf("Guess a letter or the word: ")

    // take the input
		str := ""
		fmt.Scanln(&str)


	//check if user guess the word directly
		if len(str) == len(word){
			if str == word {
				fmt.Println("You Win! ")
				break
			} else {
				entries[str] = true;
				chances -= 1
				continue
			}
		}
    // compare and update entries, placeholder and chances.
		_, ok := entries[str]
		if ok {
			//Key already exist
			continue
		}

		//update entries
		entries[str] = true;

		var found bool = false;
		for i, v := range word {
			if str == string(v) {
				placeholder[i] = string(v)
				found = true
			} 
		}
		if !found {
			chances -= 1
		}
  }
}