package main

import (
	"errors"
	"firstdemo/templates"
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
)

func main() {

	// create validate variable that not accepting number as first letter in the string

	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("must have more than 3 characters")
		}
		return nil
	}

	githubBranch := promptui.Prompt{
		Label:    "Branch",
		Validate: validate,
		Default:  "master",
	}

	dockerImageName := promptui.Prompt{
		Label:    "Docker Image Name",
		Validate: validate,
		Default:  "my-app",
	}

	port := promptui.Prompt{
		Label:   "Port Number",
		Default: "8080",
	}

	region := promptui.Select{
		Label: "Select Region",
		// Default: "asia-southeast1-a",
		Items: []string{"asia-east1-a", "asia-east1-b", "asia-east1-c", "asia-east2-a", "asia-east2-b", "asia-east2-c",
			"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c", "asia-northeast2-a", "asia-northeast2-b", "asia-northeast2-c",
			"asia-northeast3-a", "asia-northeast3-b", "asia-northeast3-c", "asia-south1-a", "asia-south1-b", "asia-south1-c",
			"asia-south2-a", "asia-south2-b", "asia-south2-c", "*asia-southeast1-a", "asia-southeast1-b", "asia-southeast1-c",
			"asia-southeast2-a", "asia-southeast2-b", "asia-southeast2-c", "australia-southeast1-a", "australia-southeast1-b", "australia-southeast1-c",
			"australia-southeast2-a", "australia-southeast2-b", "australia-southeast2-c", "europe-central2-a", "europe-central2-b", "europe-central2-c",
			"europe-north1-a", "europe-north1-b", "europe-north1-c", "europe-west1-b", "europe-west1-c", "europe-west1-d", "europe-west2-a",
			"europe-west2-b", "europe-west2-c", "europe-west3-a", "europe-west3-b", "europe-west3-c", "europe-west4-a", "europe-west4-b", "europe-west4-c",
			"europe-west6-a", "europe-west6-b", "europe-west6-c", "northamerica-northeast1-a", "northamerica-northeast1-b", "northamerica-northeast1-c",
			"northamerica-northeast2-a", "northamerica-northeast2-b", "northamerica-northeast2-c", "southamerica-east1-a", "southamerica-east1-b", "southamerica-east1-c",
			"southamerica-west1-a", "southamerica-west1-b", "southamerica-west1-c", "us-central1-a", "us-central1-b", "us-central1-c", "us-central1-f",
			"us-east1-b", "us-east1-c", "us-east1-d", "us-east4-a", "us-east4-b", "us-east4-c", "us-east4-d", "us-west1-a", "us-west1-b", "us-west1-c",
			"us-west2-a", "us-west2-b", "us-west2-c", "us-west3-a", "us-west3-b", "us-west3-c", "us-west4-a", "us-west4-b", "us-west4-c",
		},
	}

	firebaseAuth := promptui.Select{
		Label: "Select Day",
		Items: []string{"Yes", "No"},
	}

	answer1, err := githubBranch.Run()
	answer2, err := dockerImageName.Run()
	answer3, err := port.Run()
	_, answer4, err := region.Run()
	_, answer5, err := firebaseAuth.Run()

	// regex to get rid of *symbol
	reg := regexp.MustCompile(`\*`)
	answer4_final := reg.ReplaceAllString(answer4, "${1}")

	if answer5 == "Yes" {
		answer5 = "allow-authenticated"
	} else {
		answer5 = "allow-unauthenticated"
	}

	if err != nil {
		fmt.Printf("Opps, something wong with the tools, please try again. %v\n", err)
		return
	}

	// add this answer to the template && save this as cloud-run-action.yaml on the same directory
	fmt.Printf("%s\n", templates.Gaction(answer1, answer2, answer3, answer4_final, answer5))
	// fmt.Printf("answer1 %q\n", answer1)
	// fmt.Printf("answer2 %q\n", answer2)
	// fmt.Printf("answer3 %q\n", answer3)
	// fmt.Printf("answer4 %q\n", answer4_final)
	// fmt.Printf("answer4 %q\n", answer5)
}
