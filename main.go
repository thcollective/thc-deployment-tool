package main

import (
	"errors"
	"firstdemo/templates"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/TwiN/go-color"
	"github.com/manifoldco/promptui"
)

func main() {

	// initialize variables
	globalPort := ""

	// validation length
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("must have more than 3 characters")
		}
		return nil
	}

	/* START DOCKERFILE FILE CREATION*/
	fmt.Println(color.Bold + "PHASE: DOCKERFILE CREATION" + color.Reset)

	project := promptui.Select{
		Label: "What project are you working on?",
		Items: []string{"Frontend", "Backend"},
	}
	_, ansProject, _ := project.Run()

	if ansProject == "Frontend" {
		framework := promptui.Select{
			Label: "What framework are you using?",
			Items: []string{"Vue", "Nuxt"},
		}

		_, ansFramework, _ := framework.Run()
		if ansFramework == "Vue" {
			// default to port 80 for nginx integration
			portNo := "80"
			globalPort = portNo
			f, err := os.Create("Dockerfile")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			val := templates.Vuedocker(portNo)
			data := []byte(val)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

			folderPath := "nginx"
			os.MkdirAll(folderPath, os.ModePerm)

			f2, err3 := os.Create("nginx/nginx.conf")

			if err3 != nil {
				log.Fatal(err3)
			}

			defer f2.Close()

			val2 := templates.Nginx(portNo)
			data2 := []byte(val2)

			_, err4 := f2.Write(data2)

			if err4 != nil {
				log.Fatal(err4)
			}

		} else if ansFramework == "Nuxt" {

			portQuest := promptui.Prompt{
				Label:   "What port number did you exposed on your dockerfile?",
				Default: "3000",
			}

			portSelected, _ := portQuest.Run()

			globalPort = portSelected

			f, err := os.Create("Dockerfile")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			val := templates.Nuxtdocker(portSelected)
			data := []byte(val)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

		}
	}

	if ansProject == "Backend" {
		framework := promptui.Select{
			Label: "What framework are you using?",
			Items: []string{"ExpressJS", "goFiber"},
		}

		_, ansFramework, _ := framework.Run()
		if ansFramework == "ExpressJS" {

			portQuest := promptui.Prompt{
				Label:   "What port number did you exposed on your dockerfile?",
				Default: "5000",
			}

			portSelected, _ := portQuest.Run()

			globalPort = portSelected

			f, err := os.Create("Dockerfile")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			val := templates.Expressdocker(portSelected)
			data := []byte(val)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

			// will create nginx folder also

		} else if ansFramework == "goFiber" {

			portQuest := promptui.Prompt{
				Label:   "What port number did you exposed on your dockerfile?",
				Default: "5000",
			}

			portSelected, _ := portQuest.Run()

			globalPort = portSelected

			f, err := os.Create("Dockerfile")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			val := templates.Godocker(portSelected)
			data := []byte(val)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

		}
	}
	/* END DOCKERFILE FILE CREATION*/

	/* START SONARCLOUD GITHUB ACTIONS*/
	fmt.Println(color.Bold + "PHASE: SONARCLOUD ACTIONS" + color.Reset)

	sonarRootDir, sonarTestDir, sonarTestInclusions, sonarTestExclusions := "", "", "", ""

	fSonarProps, fSonarPropsErr := os.Create("sonar-project.properties")

	if fSonarPropsErr != nil {
		log.Fatal(fSonarProps)
	}

	defer fSonarProps.Close()

	orgKey := promptui.Prompt{
		Label:   "Add your sonarcloud organization key",
		Default: "thcollective",
	}

	projKey := promptui.Prompt{
		Label:   "Add your sonarcloud project key",
		Default: "",
	}

	sonarIgnore := promptui.Select{
		Label: "Do you want to ignore specific directories / files from sonarcloud?",
		Items: []string{"*Yes", "No"},
	}

	_, ansIgnore, _ := sonarIgnore.Run()

	if ansIgnore == "*Yes" {
		rootDir := promptui.Prompt{
			Label:   "Please specify the root directory of your project",
			Default: "",
		}

		testDirectory := promptui.Select{
			Label: "Do you have a test directory?",
			Items: []string{"Yes", "No"},
		}

		_, testDir, _ := testDirectory.Run()
		rootDirSelected, _ := rootDir.Run()
		sonarRootDir = rootDirSelected

		if testDir == "Yes" {
			fmt.Println("Assuming your test directory is the same as your root directory at " + color.Bold + sonarRootDir + color.Reset)
			sonarTestDir = "sonar.tests = " + sonarRootDir

			testInclusions := promptui.Select{
				Label: "Do you want to include test subdirectories in the test scope?",
				Items: []string{"Yes", "No"},
			}

			testExclusions := promptui.Select{
				Label: "Do you want to exclude test subdirectories from the test scope?",
				Items: []string{"Yes", "No"},
			}

			_, testInclusionsSelected, _ := testInclusions.Run()
			_, testExclusionsSelected, _ := testExclusions.Run()

			if testInclusionsSelected == "Yes" {
				sonarInclusions := "sonar.test.inclusions = src/**/test/**/*"
				sonarTestInclusions = sonarInclusions
			}

			if testExclusionsSelected == "Yes" {
				sonarExclusions := "sonar.test.exclusions = src/**/test/**/*"
				sonarTestExclusions = sonarExclusions
			}
		}

	}

	sonarOrgKey, _ := orgKey.Run()
	sonarProjKey, _ := projKey.Run()

	valSonarProps := templates.SonarProps(sonarOrgKey, sonarProjKey, sonarRootDir, sonarTestDir, sonarTestInclusions, sonarTestExclusions)
	dataSonarProps := []byte(valSonarProps)

	_, errSonarProps := fSonarProps.Write(dataSonarProps)

	if errSonarProps != nil {
		log.Fatal(errSonarProps)
	}

	// scannerBranch := promptui.Prompt{
	// 	Label:   "Which branch would you like to run the code scanner?",
	// 	Default: "main",
	// }

	// sonarBranch, _ := scannerBranch.Run()

	// folderPathSonar := ".github/workflows"
	// os.MkdirAll(folderPathSonar, os.ModePerm)
	// fSonarcloud, fSonarErr := os.Create(".github/workflows/sonarcloud.yaml")

	// if fSonarErr != nil {
	// 	log.Fatal(fSonarErr)
	// }

	// defer fSonarcloud.Close()

	// valSonar := templates.Sonaraction(sonarBranch)
	// dataSonar := []byte(valSonar)

	// _, errSonar := fSonarcloud.Write(dataSonar)

	// if errSonar != nil {
	// 	log.Fatal(errSonar)
	// }

	/* END SONARCLOUD GITHUB ACTIONS*/

	/* START SENTRY GITHUB ACTIONS*/
	// fmt.Println("PHASE: SENTRY GITHUB ACTIONS FILE CREATION")

	// sentry := promptui.Select{
	// 	Label: "Do you want to use sentry? (only for frontend)",
	// 	Items: []string{"Yes", "No"},
	// }

	// _, ansSentry, _ := sentry.Run()

	// if ansSentry == "Yes" {

	// 	folderPathSentry := ".github/workflows"
	// 	os.MkdirAll(folderPathSentry, os.ModePerm)
	// 	fSentry, fSenErr := os.Create(".github/workflows/sentry.yaml")

	// 	if fSenErr != nil {
	// 		log.Fatal(fSenErr)
	// 	}

	// 	defer fSentry.Close()

	// 	valSentry := templates.Sentryaction()
	// 	dataSentry := []byte(valSentry)

	// 	_, errSentry := fSentry.Write(dataSentry)

	// 	if errSentry != nil {
	// 		log.Fatal(errSentry)
	// 	}
	// }

	/* END SENTRY GITHUB ACTIONS*/

	// TODO github action for CI/CD -> for which actions runs first
	// assignees: ass77

	/* START CLOUD RUN GITHUB ACTIONS*/
	fmt.Println(color.Bold + "PHASE: CLOUD RUN ACTIONS" + color.Reset)

	purpose := promptui.Select{
		Label: "Please select your purpose for creating this cloud run action files",
		Items: []string{"for development with development environment (main)", "for development with production environment (staging)", "for production with production environment (production)"},
	}

	_, branching, _ := purpose.Run()
	answer1 := ""

	if branching == "for development with development environment (main)" {
		answer1 = "main"
	} else if branching == "for development with production environment (staging)" {
		answer1 = "staging"
	} else if branching == "for production with production environment (production)" {
		answer1 = "production"
	}

	dockerImageName := promptui.Prompt{
		Label:    "What is the name of your docker image?",
		Validate: validate,
		Default:  "my-app",
	}

	port := promptui.Prompt{
		Label:   "What port number did you exposed on your dockerfile?",
		Default: globalPort,
	}

	region := promptui.Select{
		Label: "Please select region that you want to deploy to",
		Items: []string{"*asia-southeast1", "asia-east1-a", "asia-east1-b", "asia-east1-c", "asia-east2-a", "asia-east2-b", "asia-east2-c",
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

	// answer1, _ := githubBranch.Run()
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

	// env := promptui.Select{
	// 	Label: "Do you need to set environment variables?",
	// 	Items: []string{"Yes", "No"},
	// }

	// _, answer8, err := env.Run()

	// TODO setting up environment variables in cloud run (need array | object)
	// assignees: ass77

	// envNames := ""
	// envValues := ""``
	// answerEnv := ""

	// if answer8 == "Yes" {

	// 	setEnvNames := promptui.Prompt{
	// 		Label:   "Set your environment name",
	// 		Default: "",
	// 	}

	// 	setEnvValues := promptui.Prompt{
	// 		Label:   "Set your environment value",
	// 		Default: "",
	// 	}

	// 	answer9 := ""
	// 	answer9, _ = setEnvNames.Run()
	// 	envNames = answer9

	// 	answer10 := ""
	// 	answer10, _ = setEnvValues.Run()
	// 	envValues = answer10

	// 	more := promptui.Select{
	// 		Label: "Do you need to add more environment variables?",
	// 		Items: []string{"Yes", "No"},
	// 	}

	// 	_, answer11, _ := more.Run()

	// 	if answer11 == "Yes" {

	// 		fmt.Printf("more env...")
	// 		// repeat setEnvNames and setEnvValues and push value to the array
	// 	}

	// 	// answerEnv will print x times (based on the array length) and loop the envNames and envValues (different values)
	// 	answerEnv = fmt.Sprintf(`--set-env-vars %s=%s`, envNames, envValues)

	// }

	// if err != nil {
	// 	fmt.Printf("Opps, something wong with the tools, please try again. %v\n", err)
	// 	return
	// }

	// folderPathGaction := ".github/workflows"
	// os.MkdirAll(folderPathGaction, os.ModePerm)

	// fGaction, errGaction := os.Create(".github/workflows/cloud-run-action.yaml")

	// if errGaction != nil {
	// 	log.Fatal(errGaction)
	// }

	// defer fGaction.Close()

	// valGaction := templates.Gaction(answer1, answer2, answer3, answer4_final, answer5)
	// dataGaction := []byte(valGaction)

	// _, err2 := fGaction.Write(dataGaction)

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	/* END CLOUD RUN GITHUB ACTIONS*/

	/* START TD TO ISSUE ACTIONS */
	fmt.Println(color.Bold + "PHASE: TODO TO ISSUE ACTIONS CREATION" + color.Reset)

	todo := promptui.Select{
		Label: "Do you want to add TODO to Issue github actions? ",
		Items: []string{"Yes", "No"},
	}

	_, includeTodo, _ := todo.Run()
	answerTodo := ""

	if includeTodo == "Yes" {
		// folderPathTodo := ".github/workflows"
		// os.MkdirAll(folderPathTodo, os.ModePerm)
		// fTodo, errTodo := os.Create(".github/workflows/todo-issue.yaml")

		// if errTodo != nil {
		// 	log.Fatal(errTodo)
		// }

		// defer fTodo.Close()

		// valTodo := templates.Todoaction()
		// dataTodo := []byte(valTodo)

		// _, errTodo2 := fTodo.Write(dataTodo)

		// if errTodo2 != nil {
		// 	log.Fatal(errTodo2)
		// }

		answerTodo = `  todo:
	    runs-on: ubuntu-latest
	    steps:
	    - uses: actions/checkout@master
	    - name: TODO to Issue
	        uses: alstr/todo-to-issue-action@v4.5
	        id: todo`
	}

	fmt.Printf(color.Red + "\nYou might need to fix the indentation issues in generated .yaml file later on.\n" + color.Reset)

	/* END TD TO ISSUE ACTIONS */

	folderPathAll := ".github/workflows"
	os.MkdirAll(folderPathAll, os.ModePerm)
	fAll, fAllErr := os.Create(".github/workflows/thc-deployment.yaml")

	if fAllErr != nil {
		log.Fatal(fAllErr)
	}

	defer fAll.Close()

	valAll := templates.ThcToolKit(answer1, answer2, answer3, answer4_final, answer5, answerTodo)
	dataAll := []byte(valAll)

	_, errAll2 := fAll.Write(dataAll)

	if errAll2 != nil {
		log.Fatal(errAll2)
	}

	// fmt.Println(color.Bold + "PHASE: COMMITLINT ACTIONS" + color.Reset)
	// /* START COMMITLINT ACTIONS */

	// /* END COMMITLINT ACTIONS */

	// fmt.Println(color.Bold + "PHASE: SEMANTIC RELEASES ACTIONS" + color.Reset)

	// /* START SEMANTIC RELEASES ACTIONS */

	// /* END SEMANTIC RELEASES ACTIONS */

	fmt.Printf("\n" + color.Green + "THCFileSystem " + color.Reset + "has successfully been created. Push the repo to your " + color.Red + "branch" + color.Reset + " and you're good to go!\n")
	fmt.Printf("\np/s: Please reach out to" + color.Blue + " Adri or Ming " + color.Reset + "for the" + color.Yellow + " secrets " + color.Reset + "before you make a commit.\n\n")

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
