## Procedure

a. Adri or Ming will provide the secrets to the user personally

```
GCP_PROJECT_ID ---> project_id
GCP_SA_EMAIl --> github action deploy email with roles
GCP_SA_KEY ---> GCP_SA_EMAIL secret key (.json)
SONAR_TOKEN ---> sonar token for github action (organizations)
SONAR_ORG ---> sonar properties
SONAR_PROJECT_KEY --> sonar properties
```

`NOTE`: make sure in the GCP projects, all API services + IAM roles are set based on project requirements

The default API services are as follow:

![gcloud services](/img/gcloud-services.png)


The default IAM roles are as follow (iam.gserviceaccount.com):

1. Cloud Run Admin
2. Service Account User
3. Storage Admin


### To generate new template

1. Go to `/templates` and create new "name".qtpl file
2. [Template generator reference](https://github.com/valyala/quicktemplate)
3. To compile qtc, run `qtc "name".qtpl` or `qtc .`

### Future Work

- [x] TODO to Issue Action
- [x] Commit Lint on pull request
- [x] Semantic Releases
- [x] Automated CI tests with Postman and Newman - applicable for BE only
- [x] Allow user to set environment variables on Dockerfile and cloud run action yaml file (current status: onlyset one env var @ cloud run action)
- [/] Github authentication to run the script (thc-token)
- [/] Create new github actions file to determine which actions to run first


### How To Push to Releases

1. git tag -m "Release vx.x.x" vx.x.x HEAD
2. git push --tags

<p align="right">(<a href="#top">back to top</a>)</p>