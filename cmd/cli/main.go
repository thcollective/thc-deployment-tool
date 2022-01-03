package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"thcdeploymentool/templates"

	"github.com/TwiN/go-color"
	"github.com/manifoldco/promptui"
)

func main() {

	// initialize global variables
	globalPort := ""
	ansEnvFile := ""

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

			mainDir := promptui.Prompt{
				Label:   "Please specify where is your main.go located. ",
				Default: "cmd/app/main.go",
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
				Label:   "Please speficy the folder name that you want to test",
				Default: "test",
			}

			ApiFolderSelected, _ := ApiFolder.Run()

			ApiFiles := promptui.Prompt{
				Label:   "Please speficy the collection name that you want to test",
				Default: "collection.json",
			}

			ApiFilesSelected, _ := ApiFiles.Run()

			ApiEnv := promptui.Prompt{
				Label:   "Please speficy the environment file for testing",
				Default: "environment.json",
			}

			ApiEnvSelected, _ := ApiEnv.Run()

			testAPIBranch := promptui.Prompt{
				Label:   "Which branch would you like to run the api testing github actions?",
				Default: "development, staging",
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

	/* END DOCKERFILE FILE CREATION*/

	/* START SONARCLOUD GITHUB ACTIONS*/
	fmt.Println(color.Bold + "PHASE: SONARCLOUD ACTIONS" + color.Reset)

	runSonar := promptui.Select{
		Label: "Do you want to run sonarcloud actions?",
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

		sonarRootDir, sonarTestDir, sonarTestInclusions, sonarTestExclusions := "", "", "", ""

		fSonarProps, fSonarPropsErr := os.Create("sonar-project.properties")

		if fSonarPropsErr != nil {
			log.Fatal(fSonarProps)
		}

		defer fSonarProps.Close()

		orgKey := promptui.Prompt{
			Label:   "Add your sonarcloud organization key",
			Default: "${{secrets.SONAR_ORG}}",
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

			rootDirSelected, _ := rootDir.Run()

			testDirectory := promptui.Select{
				Label: "Do you have a test directory to scan? If so, is your test code is intermingled with your source code?",
				Items: []string{"Yes", "No", "I don't have any test directory"},
			}
			_, testDir, _ := testDirectory.Run()
			sonarRootDir = rootDirSelected

			if testDir == "Yes" {

				sonarTestDir = "sonar.tests = " + rootDirSelected

				fmt.Println(color.Blue + `You may need to modify the values for sonar.test.inclusions in sonar-project.properties file based on your own test directory-level` + color.Reset)

				sonarInclusions := "sonar.test.inclusions = " + rootDirSelected + "/**/test/**/*"
				sonarTestInclusions = sonarInclusions

				sonarExclusions := "sonar.exclusions = " + rootDirSelected + "/**/test/**/*"
				sonarTestExclusions = sonarExclusions

			} else {
				testDirName := promptui.Prompt{
					Label:   "Please specify the test directory in the project",
					Default: "tests",
				}

				testDirNameSelected, _ := testDirName.Run()

				sonarTestDir = "sonar.tests = " + testDirNameSelected

				testExclude := promptui.Select{
					Label: "Do you want to exclude some folders from code scanning?",
					Items: []string{"Yes", "No"},
				}
				_, testExcludeDir, _ := testExclude.Run()

				if testExcludeDir == "Yes" {

					excludeDir := promptui.Prompt{
						Label:   "Please specify the directory that you want to exclude from code scanning",
						Default: "foo, bar",
					}

					excludeDirSelected, _ := excludeDir.Run()

					sonarExclusions := "sonar.exclusions = " + excludeDirSelected + "/**/test/**/*"
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
	}

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
		Items: []string{"for development with development environment (development)", "for development with production environment (staging)", "for production with production environment (production)"},
	}

	_, branching, _ := purpose.Run()
	answer1 := ""

	if branching == "for development with development environment (development)" {
		answer1 = "development"
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

	/* START TD TO ISSUE ACTIONS */
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

	/* END TD TO ISSUE ACTIONS */

	/* START COMMITLINT ACTIONS */
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
	/* END COMMITLINT ACTIONS */

	/* START SEMANTIC RELEASES ACTIONS */
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

	/* END SEMANTIC RELEASES ACTIONS */

	/* START thc-deployment.yaml file creation */
	fmt.Println(color.Bold + "PHASE: FINALIZING DEPLOYMENT FILE WITH ENV" + color.Reset)
	envFile := promptui.Select{
		Label: "Do you have/use .env file?",
		Items: []string{"Yes", "No"},
	}

	_, includeEnvFile, _ := envFile.Run()

	if includeEnvFile == "Yes" {

		fmt.Println(color.Yellow + "Please provide the path of your .env file" + color.Reset)

		// create an array to store environment name and value
		env := []string{}

		// prompt user to enter the environment name and value and push to env array
		for {
			envName := promptui.Prompt{
				Label: "Environment name",
			}

			envNameVal, _ := envName.Run()

			envValue := promptui.Prompt{
				Label:   "Environment value",
				Default: "${{secret.ENV_NAME}} if you using github secrets",
			}

			envValueVal, _ := envValue.Run()

			env = append(env, envNameVal+"="+envValueVal)

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
			// fmt.Println("echo " + val + " >> .env")
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

	valAll := templates.ThcToolKit(answer1, answer2, answer3, answer4_final, answer5, includeEnvFile, ansEnvFile)
	dataAll := []byte(valAll)

	_, errAll2 := fAll.Write(dataAll)

	if errAll2 != nil {
		log.Fatal(errAll2)
	}

	/* END thc-deployment.yaml file creation */

	fmt.Printf("\n" + color.Green + "THC magic " + color.Reset + "has successfully been casted. Push the repo to your " + color.Red + "branch" + color.Reset + " and you're good to go!\n")
	fmt.Printf("\np/s: Please reach out to" + color.Blue + " Adri or Ming " + color.Reset + "for the" + color.Yellow + " secrets " + color.Reset + "before you make a commit.\n")
	fmt.Printf(color.Yellow + "\nNOTE: You might need to double check fix the yaml indentation issues in generated .yaml file.\n" + color.Reset)

}

// func HtmlspecialChars(s string) string {
// 	return html.EscapeString(s)

// }

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
