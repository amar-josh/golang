// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// 	"time"
// )

// var dev = flag.Bool("dev", false, "dev mode")

// const (
// 	GAME_PLAY = iota // untyped int constants
// 	GAME_WIN
// 	GAME_LOSE
// 	GAME_TIMEOUT
// 	GAME_ERR
// )

// type Hangman struct {
// 	Entries     map[string]bool
// 	Placeholder []string
// 	chances     int
// 	word        string
// 	duration    time.Duration // game duration
// 	input       io.Reader
// }

// func get_keys(entries map[string]bool) (keys []string) {
// 	for k, _ := range entries {
// 		keys = append(keys, k)
// 	}
// 	return
// }

// func get_word() string {
// 	if *dev {
// 		return "elephant"
// 	}
// 	resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
// 	if err != nil {
// 		return "elephant"
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)

// 	var words []string
// 	err = json.Unmarshal(body, &words)
// 	if err != nil {
// 		// handle error
// 		return "elephant"
// 	}

// 	//fmt.Println(words)
// 	for _, word := range words {
// 		if len(word) > 4 && len(word) < 9 {
// 			return word
// 		}
// 	}

// 	return words[0]
// }

// func (h *Hangman) Render(status uint) {
// 	switch status {
// 	case GAME_PLAY:
// 		fmt.Println(h.Placeholder)                       // render the placeholder
// 		fmt.Printf("Chances: %d\n", h.chances)           // render the chances left
// 		fmt.Printf("Entries: %v\n", get_keys(h.Entries)) // show the letters or words guessed till now.
// 		fmt.Printf("Guess a letter or the word: ")

// 	case GAME_WIN:
// 		fmt.Println("You win!!")
// 	case GAME_LOSE:
// 		fmt.Println("You're out of chances")
// 		fmt.Println("Word was: ", h.word)
// 		fmt.Println("Game Over! Try again")

// 	case GAME_ERR:
// 		fmt.Println("Something wrong with getting input")

// 	case GAME_TIMEOUT:
// 		fmt.Println("Timedout... too bad!")
// 	}
// }

// func (h *Hangman) GetInput() (str string) {
// 	str = ""
// 	fmt.Scanln(&str)
// 	return
// }

// func get_input(input io.Reader) (str string, err error) {
// 	scanner := bufio.NewScanner(input)
// 	if scanner.Scan() { // scan one line
// 		str = scanner.Text()
// 		return
// 	}

// 	err = scanner.Err()
// 	return
// }

// func play(h *Hangman, result chan bool) {
// 	for {
// 		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
// 		userInput := strings.Join(h.Placeholder, "")
// 		if h.chances == 0 && userInput != h.word {
// 			h.Render(GAME_LOSE)
// 			result <- false
// 			return
// 		}

// 		// evaluate a win!
// 		if userInput == h.word {
// 			h.Render(GAME_WIN)
// 			result <- true
// 			return
// 		}

// 		// Console display
// 		h.Render(GAME_PLAY)
// 		// Addon validation: Allow only alpha!
// 		// Addon validation: manage case!

// 		str := h.GetInput()

// 		// if len(str) > 2, compare the word with the str
// 		if len(str) > 2 {
// 			if str == h.word {
// 				h.Render(GAME_WIN)
// 				result <- true
// 				return
// 			} else {
// 				// you lose a chance
// 				h.Entries[str] = true
// 				h.chances -= 1
// 				continue
// 			}
// 		}

// 		// compare and update entries, placeholder and chances.
// 		_, ok := h.Entries[str]
// 		if ok {
// 			// key exists already; duplicate
// 			continue
// 		}

// 		h.Entries[str] = true
// 		// check if letter exists in the word!
// 		found := false
// 		for i, v := range h.word {
// 			if str == string(v) {
// 				h.Placeholder[i] = string(v)
// 				found = true
// 			}
// 		}

// 		if !found {
// 			h.chances -= 1
// 		}
// 	}
// }

// func main() {
// 	flag.Parse()

// 	h := Hangman{
// 		Entries:     map[string]bool{},
// 		Placeholder: []string{},
// 		chances:     8,
// 		duration:    2 * time.Minute,
// 	}

// 	h.word = get_word()

// 	// get length of the word len(word)
// 	// initialize slice with each element as "_"
// 	for i := 0; i < len(h.word); i++ {
// 		h.Placeholder = append(h.Placeholder, "_")
// 		//placeholder[i] = "_"
// 	}

// 	t := time.NewTimer(h.duration)
// 	result := make(chan bool)

// 	go play(&h, result)

// 	// Wait on either game to finish or timeout!
// 	select {
// 	case <-result:
// 	case <-t.C:
// 		h.Render(GAME_TIMEOUT)
// 	}
// }