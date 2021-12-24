package main

import (
	"errors"
	"firstdemo/templates"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
)

func main() {

	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("must have more than 3 characters")
		}
		return nil
	}

	// TODO: DOCKERFILE

	// TODO: SONARCLOUD GITHUB ACTIONS

	// TODO: SENTRY GITHUB ACTIONS

	// CLOUD RUN GITHUB ACTIONS
	githubBranch := promptui.Prompt{
		Label:    "What is your primary branch? Is it main or master?",
		Validate: validate,
		Default:  "master",
	}

	dockerImageName := promptui.Prompt{
		Label:    "What project are you working on? *must be same as docker image name",
		Validate: validate,
		Default:  "my-app",
	}

	port := promptui.Prompt{
		Label:   "What port number do you expose on your dockerfile?",
		Default: "8080",
	}

	region := promptui.Select{
		Label: "Please select region that you want to deploy to",
		Items: []string{"*asia-southeast1-a", "asia-east1-a", "asia-east1-b", "asia-east1-c", "asia-east2-a", "asia-east2-b", "asia-east2-c",
			"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c", "asia-northeast2-a", "asia-northeast2-b", "asia-northeast2-c",
			"asia-northeast3-a", "asia-northeast3-b", "asia-northeast3-c", "asia-south1-a", "asia-south1-b", "asia-south1-c",
			"asia-south2-a", "asia-south2-b", "asia-south2-c", "asia-southeast1-b", "asia-southeast1-c",
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
		Label: "Are you using Firebase Authentication Service?",
		Items: []string{"Yes", "No"},
	}

	// authenticating service to service
	s2s := promptui.Select{
		Label: "Is your architecture using multiple services ? *authenticating service-to-service",
		Items: []string{"Yes", "No"},
	}

	// TODO: will call GCP api to create new project | use existing project + IAM roles automation + API services activation automation
	// webhooks ?
	answer1, _ := githubBranch.Run()
	answer2, _ := dockerImageName.Run()
	answer3, _ := port.Run()
	_, answer4, _ := region.Run()
	_, answer5, _ := firebaseAuth.Run()
	_, answer6, _ := s2s.Run()

	// regex to get rid of *symbol
	reg := regexp.MustCompile(`\*`)
	answer4_final := reg.ReplaceAllString(answer4, "${1}")

	if answer5 == "Yes" {
		answer5 = "allow-authenticated"
	} else {
		answer5 = "allow-unauthenticated"
	}

	if answer6 == "Yes" {
		lang := promptui.Select{
			Label: "What language are you using?",
			Items: []string{
				"Javascript", "Python", "Go",
			},
		}
		_, answer7, _ := lang.Run()

		extension := ""
		// file extension
		if answer7 == "Go" {
			extension = "go"
		} else if answer7 == "Javascript" {
			extension = "js"
		} else if answer7 == "Python" {
			extension = "py"
		}

		filename := fmt.Sprintf("services.%s", extension)
		fx, err := os.Create(filename)

		if err != nil {
			log.Fatal(err)
		}

		if answer7 == "Javascript" {
			answer7 = "js"
			file := templates.JavascriptS2S(answer7)
			data := []byte(file)

			_, err2 := fx.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

		}
		if answer7 == "Python" {
			answer7 = "py"
			file := templates.PythonS2S(answer7)
			data := []byte(file)

			_, err2 := fx.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

		}
		if answer7 == "Go" {
			answer7 = "go"
			file := templates.GoS2S(answer7)
			data := []byte(file)
			_, err2 := fx.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

		}

		defer fx.Close()

	}

	// TODO: prompts for app environment variables
	env := promptui.Select{
		Label: "Do you need to set environment variables?",
		Items: []string{"Yes", "No"},
	}

	_, answer8, err := env.Run()

	// TODO: these needs to be an array if env > 1
	envNames := ""
	envValues := ""
	answerEnv := ""

	// TODO: make it in array if answer11 == "Yes", then loop?
	if answer8 == "Yes" {

		setEnvNames := promptui.Prompt{
			Label:   "Set your environment name",
			Default: "",
		}

		setEnvValues := promptui.Prompt{
			Label:   "Set your environment value",
			Default: "",
		}

		answer9 := ""
		answer9, _ = setEnvNames.Run()
		envNames = answer9

		answer10 := ""
		answer10, _ = setEnvValues.Run()
		envValues = answer10

		more := promptui.Select{
			Label: "Do you need to add more environment variables?",
			Items: []string{"Yes", "No"},
		}

		_, answer11, _ := more.Run()

		if answer11 == "Yes" {

			fmt.Printf("more env...")
			// repeat setEnvNames and setEnvValues and push value to the array
		}

		// answerEnv will print x times (based on the array length) and loop the envNames and envValues (different values)
		answerEnv = fmt.Sprintf(`--set-env-vars %s=%s`, envNames, envValues)

	}

	if err != nil {
		fmt.Printf("Opps, something wong with the tools, please try again. %v\n", err)
		return
	}

	finalPrompt := promptui.Select{
		Label: "All good? ",
		Items: []string{"Yes", "No"},
	}

	_, good, _ := finalPrompt.Run()

	// create .github/workflows folder if not exist
	folderPath := ".github/workflows"
	os.MkdirAll(folderPath, os.ModePerm)

	f, err := os.Create(".github/workflows/cloud-run-action.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if good == "Yes" {

		val := templates.Gaction(answer1, answer2, answer3, answer4_final, answer5, answerEnv)
		data := []byte(val)

		_, err2 := f.Write(data)

		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Printf("\nDeployments file has successfully been created. Push the repo to your primary branch and you're good to go!\n")

		fmt.Printf("\np/s: Please reach out to Adri or Ming for the secrets before you commit your code to your primary branch, thank you!\n")
	}

}

// func ensureDir(dirName string) error {
// 	err := os.Mkdir(dirName, os.ModeDir)
// 	if err == nil {
// 		return nil
// 	}
// 	if os.IsExist(err) {
// 		// check that the existing path is a directory
// 		info, err := os.Stat(dirName)
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() {
// 			return errors.New("path exists but is not a directory")
// 		}
// 		return nil
// 	}
// 	return err
// }
