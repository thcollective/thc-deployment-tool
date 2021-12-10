# THC Deployment Automation tools

## THe flow of this tool - v0.1

a. only organization controls the secret

```
GCP_PROJECT_ID ---> thcv1-blabla
GCP_SA_EMAIl --> githubaction email
GCP_SA_KEY ---> secret key
```
make sure in the **GCP projects, all API services + IAM roles are set


b. Prompt to generate cloud run yaml file

1. which branch do you want to perform this actions ?
	branch	
2. what is your docker image name ? 
	dockerImageName
3. which port did you expose in your Dockerfile ?
	port
4. Select which region do you want to host your app
	Region
	select Option --> like jalapeno promptui https://cloud.google.com/compute/docs/regions-zones

5. Are you using Firebase authentication?
	Select Option firebaseAuth
		Yes : allow-authenticate
		No : allow-unauthenticate


6. take the variables value and insert to the template (cloud-run-action.yaml)

7. save this on root of your project in .github/workflows folder and name it cloud-run-action.yaml

prerequisite: Need to have Dockerfile

### Future Work

1. save the template to file and called it cloud-run-action.yaml so the user can just c&p the file directly 
2. authentication to run the script (thc-token) 
