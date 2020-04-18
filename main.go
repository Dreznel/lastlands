package main

import (
	"fmt"

	"github.com/dreznel/lastlands/v2/internal/character"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
// var qs = []*survey.Question{
// 	{
// 		Name:      "name",
// 		Prompt:    &survey.Input{Message: "What is your name?"},
// 		Validate:  survey.Required,
// 		Transform: survey.Title,
// 	},
// 	{
// 		Name: "color",
// 		Prompt: &survey.Select{
// 			Message: "Choose a color:",
// 			Options: []string{"red", "blue", "green"},
// 			Default: "red",
// 		},
// 	},
// 	{
// 		Name:   "age",
// 		Prompt: &survey.Input{Message: "How old are you?"},
// 	},
// }

func main() {
	for {

		mainMenuSelection := ""
		prompt := &survey.Select{
			Message: "Welcome to the Lastlands Character Mangaer! Select an option:",
			Options: []string{"Create Character", "Manager Character"},
		}
		err := survey.AskOne(prompt, &mainMenuSelection)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch mainMenuSelection {
		case "Create Character":
			fmt.Println("Creating character...")
			createCharacterInterface()
			break
		case "Manage Character":
			fmt.Println("Sorry! The Manage Character feature isn't implemented yet!")
			break
		default:
			break
		}
	}

	// // the answers will be written to this struct
	// answers := struct {
	// 	Name          string // survey will match the question and field names
	// 	FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
	// 	Age           int    // if the types don't match, survey will convert it
	// }{}

	// // perform the questions
	// err := survey.Ask(qs, &answers)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	//fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)
}

func createCharacterInterface() {
	seeker := character.CreateCharacter()
	fmt.Println(fmt.Sprintf("They call me %s, here", seeker.Name))
	fmt.Println("----")
	fmt.Println(fmt.Sprintf("Core Memories: %v", seeker.CoreAttributes))
	fmt.Println(fmt.Sprintf("Base stats:\n %v", seeker.AbilityScores))
}
