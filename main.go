package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fmt.Println("Welcome to the Advanced CLI Wizard!")

	// Create a struct to hold user responses
	answers := struct {
		Name    string
		Age     int
		Confirm bool
		Choice  string
	}{}

	// Define the initial set of questions
	initialQuestions := []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What's your name:"},
			Validate: survey.Required,
		},
		{
			Name:     "age",
			Prompt:   &survey.Input{Message: "How old are you:"},
			Validate: validateAge,
		},
		{
			Name:   "confirm",
			Prompt: &survey.Confirm{Message: "Is this information correct:"},
		},
	}

	// Define additional questions based on the user's choice
	conditionalQuestions := []*survey.Question{
		{
			Name:   "choice",
			Prompt: &survey.Select{Message: "Select an option:", Options: []string{"Option A", "Option B"}},
		},
	}

	// Ask the initial questions
	err := survey.Ask(initialQuestions, &answers)
	if err != nil {
		log.Fatalf("Error asking questions: %v", err)
	}

	// If the user confirms, ask the conditional questions
	if answers.Confirm {
		fmt.Printf("Hello, %s! You are %d years old.\n", answers.Name, answers.Age)

		err := survey.Ask(conditionalQuestions, &answers)
		if err != nil {
			log.Fatalf("Error asking questions: %v", err)
		}

		fmt.Printf("You selected: %s\n", answers.Choice)
	} else {
		fmt.Println("Cancelled. Please restart the wizard if you'd like to try again.")
	}
}

func validateAge(val interface{}) error {
	strAge, ok := val.(string)
	if !ok {
		return fmt.Errorf("Age should be a number")
	}

	age, err := strconv.Atoi(strAge)
	if err != nil {
		return fmt.Errorf("Invalid age format")
	}

	if age < 1 {
		return fmt.Errorf("Age must be greater than or equal to 1")
	}

	return nil
};