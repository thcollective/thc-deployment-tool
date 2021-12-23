# THC Deployment Automation tool

### Procedure

Prerequisite: Need to have Dockerfile

a. Adri or Ming will provide the secrets to the user personally

```
GCP_PROJECT_ID ---> project_id
GCP_SA_EMAIl --> github action email
GCP_SA_KEY ---> IAM secret key
```

NOTE: make sure in the GCP projects, all API services + IAM roles are set based on project requirements

The default API services are as follow:

![gcloud services](https://github.com/ass77/deployment-automation-tool/blob/main/img/gcloud-services.png)


The default IAM roles are as follow (iam.gserviceaccount.com):

1. Cloud Run Admin
2. Service Account User
3. Storage Admin


b. Prompt to generate cloud run yaml file


1. which branch do you want to perform this actions ?

2. what is your docker image name ? 

3. which port did you expose in your Dockerfile ?

4. Select which region do you want to host your app https://cloud.google.com/compute/docs/regions-zones

5. Are you using Firebase authentication?
```
Yes : allow-authenticate
No : allow-unauthenticate
```
6. generate services file if the user architecture using multiple services



### Future Work

1. authentication to run the script (thc-token) 

### How To push to releases

1. git push 
2. git tag -m "description of release" vx.x.x HEAD
3. git push --tags
