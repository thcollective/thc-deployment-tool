package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"thcdeploymentool/templates"

	"github.com/TwiN/go-color"
	"github.com/manifoldco/promptui"
)

// global variables

var branching = ""
var answer2 = ""
var onlyBE = ""

// var answer3 = ""
var answer4 = ""
var answer4_final = ""
var answer5 = ""
var answer6 = ""
var globalPort = ""
var ansEnvFile = ""
var cloudRunBranching = ""

func main() {

	// BUG: not working - break script if user hits ctrl+c | d
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println('\n', sig)
			os.Exit(1)
		}
	}()

	dockerfileCreation()

	sonarcloudActions()

	cloudrunActions()

	cloudrunActionsFinal(cloudRunBranching, answer2, answer4_final, answer5, ansEnvFile)

	todotoissueActions()

	commitlintActions()

	semanticreleaseActions()

	done()

	// github action for CI/CD -> for which actions runs first

}

func dockerfileCreation() {

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
		onlyBE = ansProject
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

			mainDir := promptui.Prompt{
				Label:   "Please specify the path to your main.go file (for eg: cmd/app/main.go)",
				Default: "",
			}

			mainScriptDir, _ := mainDir.Run()

			f, err := os.Create("Dockerfile")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			val := templates.Godocker(portSelected, mainScriptDir)
			data := []byte(val)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

		}

		testApi := promptui.Select{
			Label: "Do you want to test your API automatically using Github Actions?",
			Items: []string{"Yes", "No"},
		}

		_, includeTestApi, _ := testApi.Run()

		if includeTestApi == "Yes" {

			ApiFolder := promptui.Prompt{
				Label:   "Please speficy the folder name that you want to test (for eg: testFolder)",
				Default: "",
			}

			ApiFolderSelected, _ := ApiFolder.Run()

			fmt.Println(color.Yellow + "NOTE: Please ensure that you have exported your API collection from your API client and stored in " + color.Reset + color.Bold + "./" + ApiFolderSelected + "/" + color.Reset + color.Bold + " folder" + color.Reset)

			ApiFiles := promptui.Prompt{
				Label:   "Please speficy the collection filename that you exported from Insomnia/Postman (for eg: collection.json)",
				Default: "",
			}

			ApiFilesSelected, _ := ApiFiles.Run()

			ApiEnv := promptui.Prompt{
				Label:   "Please speficy the environment collection filename that you exported from Insomnia/Postman (for eg: environment.json)",
				Default: "",
			}

			ApiEnvSelected, _ := ApiEnv.Run()

			testAPIBranch := promptui.Prompt{
				Label:   "Which git branch would you like to run the api testing actions? for eg: development, staging",
				Default: "",
			}

			testBranch, _ := testAPIBranch.Run()

			folderPathTestAPI := ".github/workflows"
			os.MkdirAll(folderPathTestAPI, os.ModePerm)
			fTestAPIcloud, fTestAPIErr := os.Create(".github/workflows/test-api.yaml")

			if fTestAPIErr != nil {
				log.Fatal(fTestAPIErr)
			}

			defer fTestAPIcloud.Close()

			valTestAPI := templates.TestAPIaction(testBranch, ApiFolderSelected, ApiFilesSelected, ApiEnvSelected)
			dataTestAPI := []byte(valTestAPI)

			_, errTestAPI := fTestAPIcloud.Write(dataTestAPI)

			if errTestAPI != nil {
				log.Fatal(errTestAPI)
			}

		}

	}

	fDockerIgnore, errDockerIgnore := os.Create(".dockerignore")

	if errDockerIgnore != nil {
		log.Fatal(errDockerIgnore)
	}

	defer fDockerIgnore.Close()

	valDockerIgnore := templates.Dockerignore()
	dataDockerIgnore := []byte(valDockerIgnore)

	_, errDockerIgnoreFile := fDockerIgnore.Write(dataDockerIgnore)

	if errDockerIgnoreFile != nil {
		log.Fatal(errDockerIgnoreFile)
	}

}

func cloudrunActions() (string, string, string, string) {
	fmt.Println(color.Bold + "PHASE: CLOUD RUN ACTIONS" + color.Reset)

	// validation length
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("must have more than 3 characters")
		}
		return nil
	}

	purpose := promptui.Select{
		Label: "Please choose your purpose for creating this cloud run action file",
		Items: []string{"for development with development environments (development)", "for development with production environments (staging)", "for production with production environments (production)"},
	}

	_, branching, _ = purpose.Run()

	if branching == "for development with development environments (development)" {
		cloudRunBranching = "development"

	} else if branching == "for development with production environments (staging)" {
		cloudRunBranching = "staging"

	} else if branching == "for production with production environments (production)" {
		cloudRunBranching = "production"

	}

	dockerImageName := promptui.Prompt{
		Label:    "Please specify a name for your docker image (for eg: my-app). This name will also be used in Google Cloud Run service name",
		Validate: validate,
		Default:  "",
	}

	region := promptui.Select{
		Label: "Please choose a region where you want your application to be deployed",
		Items: []string{"*asia-southeast1", "asia-east1", "asia-east2",
			"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c", "asia-northeast2-a", "asia-northeast2-b", "asia-northeast2-c",
			"asia-northeast3", "asia-south1",
			"asia-south2", "asia-southeast", "australia-southeast1",
			"australia-southeast2", "europe-central2",
			"europe-north1", "europe-west1", "europe-west2",
			"europe-west3", "europe-west4",
			"europe-west6", "northamerica-northeast1",
			"northamerica-northeast2", "southamerica-east1",
			"southamerica-west1", "us-central1",
			"us-east1", "us-east4", "us-west1",
			"us-west2", "us-west3", "us-west4",
		},
	}

	firebaseAuth := promptui.Select{
		Label: "Are you using Firebase Authentication Service?",
		Items: []string{"Yes", "No"},
	}

	// authenticating service to service only for BE
	if onlyBE == "Backend" {
		s2s := promptui.Select{
			Label: "Is your architecture using multiple services ? *authenticating service-to-service",
			Items: []string{"Yes", "No"},
		}
		_, answer6, _ = s2s.Run()

		if answer6 == "Yes" {
			lang := promptui.Select{
				Label: "What kind of programming language are you working on right now in this project?",
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
	}

	answer2, _ = dockerImageName.Run()
	// answer3, _ = port.Run()
	_, answer4, _ = region.Run()
	_, answer5, _ = firebaseAuth.Run()

	// regex to get rid of *symbol
	reg := regexp.MustCompile(`\*`)
	answer4_final = reg.ReplaceAllString(answer4, "${1}")

	if answer5 == "Yes" {
		answer5 = "allow-authenticated"
	} else {
		answer5 = "allow-unauthenticated"
	}

	return answer2, answer4_final, answer5, answer6

}

func cloudrunActionsFinal(cloudRunBranching string, answer2 string, answer4_final string, answer5 string, ansEnvFile string) {

	fmt.Println(color.Bold + "PHASE: ADDING CUSTOM ENVIRONMENT VARIABLES INTO CLOUD RUN ACTION FILE" + color.Reset)
	envFile := promptui.Select{
		Label: "Do you have/use any environment file (.env) in your project?",
		Items: []string{"Yes", "No"},
	}

	_, includeEnvFile, _ := envFile.Run()

	if includeEnvFile == "Yes" {

		fmt.Println(color.Yellow + "Please provide your custom environments variables like how you declared in it your .env file " + color.Reset)
		fmt.Println(color.Red + "NOTE: Please make sure to store your environment keys and values on github secrets for better security " + color.Reset)

		// create an array to store environment name and value
		env := []string{}

		// prompt user to enter the environment name and value and push to env array
		for {
			envName := promptui.Prompt{
				Label:   "Environment key name which you used in your .env file",
				Default: "",
			}

			envNameVal, _ := envName.Run()

			envValue := promptui.Prompt{
				Label:   "Environment value which you stored in github secrets for eg: ${{secret.ENV_VALUE}}",
				Default: "",
			}

			envValueVal, _ := envValue.Run()

			env = append(env, envNameVal+" = "+envValueVal)

			// prompt user to continue or not
			fmt.Println(color.Yellow + "Do you want to add more environment variables?" + color.Reset)

			addMoreEnv := promptui.Select{
				Label: "Select",
				Items: []string{"Yes", "No"},
			}

			_, addMoreEnvVal, _ := addMoreEnv.Run()

			if addMoreEnvVal == "No" {
				break
			}

		}

		for _, val := range env {
			special := `>>`
			ansEnvFile += "echo " + val + " " + special + " .env\n\t\t\t\t\t\t"
		}

	}

	folderPathAll := ".github/workflows"
	os.MkdirAll(folderPathAll, os.ModePerm)
	fAll, fAllErr := os.Create(".github/workflows/thc-deployment.yaml")

	if fAllErr != nil {
		log.Fatal(fAllErr)
	}

	defer fAll.Close()

	valAll := templates.ThcToolKit(cloudRunBranching, answer2, globalPort, answer4_final, answer5, includeEnvFile, ansEnvFile)
	dataAll := []byte(valAll)

	_, errAll2 := fAll.Write(dataAll)

	if errAll2 != nil {
		log.Fatal(errAll2)
	}

}

func sonarcloudActions() {
	fmt.Println(color.Bold + "PHASE: SONARCLOUD ACTIONS" + color.Reset)

	runSonar := promptui.Select{
		Label: "Choose Yes if you wish to perform code scanning using sonarcloud (you may skip choose NO if you already have sonarcloud.yaml file)",
		Items: []string{"Yes", "No"},
	}

	_, runSonarSelected, _ := runSonar.Run()

	if runSonarSelected == "Yes" {
		scannerBranch := promptui.Prompt{
			Label:   "Which branch would you like to run the code scanner?",
			Default: "development, staging",
		}

		sonarBranch, _ := scannerBranch.Run()

		folderPathSonar := ".github/workflows"
		os.MkdirAll(folderPathSonar, os.ModePerm)
		fSonarcloud, fSonarErr := os.Create(".github/workflows/sonarcloud.yaml")

		if fSonarErr != nil {
			log.Fatal(fSonarErr)
		}

		defer fSonarcloud.Close()

		valSonar := templates.Sonaraction(sonarBranch)
		dataSonar := []byte(valSonar)

		_, errSonar := fSonarcloud.Write(dataSonar)

		if errSonar != nil {
			log.Fatal(errSonar)
		}

		sonarInclusions, sonarExclusions := "", ""

		fSonarProps, fSonarPropsErr := os.Create("sonar-project.properties")

		if fSonarPropsErr != nil {
			log.Fatal(fSonarProps)
		}

		defer fSonarProps.Close()

		fmt.Print(color.Yellow + "Please ensure that you have sonarcloud organization key (SONAR_ORG) and project key (SONAR_PROJECT_KEY) stored in your github secrets. \n" + color.Reset)

		rootDir := promptui.Prompt{
			Label:   "Please specify the root folder  in your project, where most of your source files lives. (e.g. src, scripts, cmd)",
			Default: "",
		}

		rootDirSelected, _ := rootDir.Run()

		sonarRootDir := ""

		if rootDirSelected == "src" {
			sonarRootDir = "src"
		} else {
			sonarRootDir = rootDirSelected
		}

		includeDir := promptui.Select{
			Label: "Do you have any folder that you want to include beside " + color.Red + rootDirSelected + color.Reset + "?",
			Items: []string{"Yes", "No"},
		}

		_, includeDirSelected, _ := includeDir.Run()

		if includeDirSelected == "Yes" {

			sonarInclusionsName := promptui.Prompt{
				Label:   "Please specify the folder that you want to include (for eg: scripts)",
				Default: "",
			}

			sonarInclusionsSelected, _ := sonarInclusionsName.Run()

			sonarInclusions = "sonar.inclusions= " + sonarInclusionsSelected

		} else {
			excludeDir := promptui.Select{
				Label: "Do you have any folder that you want to exclude from getting scanned?",
				Items: []string{"Yes", "No"},
			}

			_, excludeDirSelected, _ := excludeDir.Run()

			if excludeDirSelected == "Yes" {

				sonarExclusionsName := promptui.Prompt{
					Label:   "Please specify the directory/path that you want to exclude (for eg: src/test)",
					Default: "",
				}

				sonarExclusionsSelected, _ := sonarExclusionsName.Run()

				sonarExclusions = "sonar.exclusions= " + sonarExclusionsSelected
			}

		}

		// sonarOrgKey, _ := orgKey.Run()
		// sonarProjKey, _ := projKey.Run()

		valSonarProps := templates.SonarProps(sonarRootDir, sonarInclusions, sonarExclusions)
		dataSonarProps := []byte(valSonarProps)

		_, errSonarProps := fSonarProps.Write(dataSonarProps)

		if errSonarProps != nil {
			log.Fatal(errSonarProps)
		}
	}

}

func todotoissueActions() {

	fmt.Println(color.Bold + "PHASE: TODO TO ISSUE ACTIONS CREATION" + color.Reset)

	todo := promptui.Select{
		Label: "Do you want to add TODO to Issue github actions? ",
		Items: []string{"Yes", "No"},
	}

	_, includeTodo, _ := todo.Run()

	if includeTodo == "Yes" {
		folderPathTodo := ".github/workflows"
		os.MkdirAll(folderPathTodo, os.ModePerm)
		fTodo, errTodo := os.Create(".github/workflows/todo-to-issue.yaml")

		if errTodo != nil {
			log.Fatal(errTodo)
		}

		defer fTodo.Close()

		valTodo := templates.Todoaction()
		dataTodo := []byte(valTodo)

		_, errTodo2 := fTodo.Write(dataTodo)

		if errTodo2 != nil {
			log.Fatal(errTodo2)

		}

	}

}

func commitlintActions() {
	fmt.Println(color.Bold + "PHASE: COMMITLINT ACTIONS" + color.Reset)

	fmt.Printf("\n" + color.Yellow + "Only applicable on pull request " + color.Reset)

	comlint := promptui.Select{
		Label: "Do you want to add commitlint actions?",
		Items: []string{"Yes", "No"},
	}

	_, includeCommitlint, _ := comlint.Run()

	if includeCommitlint == "Yes" {

		folderPathCommitLint := ".github/workflows"
		os.MkdirAll(folderPathCommitLint, os.ModePerm)
		fCommitLint, fCommitLintErr := os.Create(".github/workflows/commitlint.yaml")
		fCommitLintConfig, fCommitLintConfigErr := os.Create("commitlint.config.js")

		if fCommitLintErr != nil {
			log.Fatal(fCommitLintErr)
		}

		if fCommitLintConfigErr != nil {
			log.Fatal(fCommitLintConfigErr)
		}

		defer fCommitLint.Close()
		defer fCommitLintConfig.Close()

		valCommitLint := templates.Commitlint()
		dataCommitLint := []byte(valCommitLint)

		valCommitLintConfig := templates.CommitlintConfig()
		dataCommitLintConfig := []byte(valCommitLintConfig)

		_, errCommitLintFile := fCommitLint.Write(dataCommitLint)
		_, errCommitLintConfigFile := fCommitLintConfig.Write(dataCommitLintConfig)

		if errCommitLintFile != nil {
			log.Fatal(errCommitLintFile)
		}

		if errCommitLintConfigFile != nil {
			log.Fatal(errCommitLintConfigFile)
		}

	}
}

func semanticreleaseActions() {

	fmt.Println(color.Bold + "PHASE: SEMANTIC RELEASES ACTIONS" + color.Reset)

	semRelease := promptui.Select{
		Label: "Do you want to add semantic release actions?",
		Items: []string{"Yes", "No"},
	}

	_, includeSemRelease, _ := semRelease.Run()

	if includeSemRelease == "Yes" {

		semBranch := promptui.Prompt{
			Label:   "Which branch would you like to run the semantic releases?",
			Default: "release/**",
		}

		semanticBranch, _ := semBranch.Run()

		folderPathSemantic := ".github/workflows"
		os.MkdirAll(folderPathSemantic, os.ModePerm)
		fSemanticcloud, fSemanticErr := os.Create(".github/workflows/semantic-releases.yaml")

		if fSemanticErr != nil {
			log.Fatal(fSemanticErr)
		}

		defer fSemanticcloud.Close()

		valSemantic := templates.Sonaraction(semanticBranch)
		dataSemantic := []byte(valSemantic)

		_, errSemantic := fSemanticcloud.Write(dataSemantic)

		if errSemantic != nil {
			log.Fatal(errSemantic)
		}

	}

}

func done() {

	fmt.Printf("\n" + color.Green + "THC magic " + color.Reset + "has successfully been casted. Push your project repo to your " + color.Red + "respective branch" + color.Reset + " and you're good to go!\n")
	fmt.Printf("\np/s: Please reach out to" + color.Blue + " Adri or Ming " + color.Reset + "for the" + color.Yellow + " secrets " + color.Reset + "before you commit.\n")
	fmt.Printf(color.Yellow + "\nNOTE: You might need to double check and fix the indentation issues on generated .yaml files.\n\n" + color.Reset)

}
