package main

import "fmt"

// START OMIT
type youKnowNothingError struct {
	guess int
}

func (e youKnowNothingError) Error() string {
	return fmt.Sprintf("%d!!! You know nothing...", e.guess)
}

func isThisTheAnswerToEverything(guess int) (string, error) {
	if guess != 42 {
		return "", youKnowNothingError{guess}
	}
	return fmt.Sprintf("Deep Thought agrees, %d is the answer", guess), nil
}

// END OMIT

// START 2OMIT
func main() {

	for _, guess := range []int{7, 42} {
		message, error := isThisTheAnswerToEverything(guess)
		if error != nil {
			fmt.Println(error.Error())
		} else {
			fmt.Println(message)
		}
	}
}

// END 2OMIT
