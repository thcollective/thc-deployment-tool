<div align="center">
  <a href="https://github.com/thcollective">
    <img src="img/thc.png" alt="thc_logo" width="80" height="80">
  </a>
  <h3 align="center">THC deployment automation tool</h3>
</div>

### What is it about?

This deployment automation tool helps users to generate basic files that are required to deploy their app on cloud run, scan their code to check for bugs before deployment and tracing any error on the app in dev mode or prod mode. The script will generate files as follows:

* Dockerfile for frontend (vue, nuxt) and backend(expressJS, gofiber)
* Sonarcloud github actions and properties file (need to seek from Admin for keys)
* Sentry github actions and properties files (under development)
* Cloud Run github actions 

### How To Run This Project (production)

1. Download the latest binary package [here](https://github.com/thcollective/thc-deployment-tool/releases)

2. Run `./main` on your working project root directory, start answering the prompt and then the spell shall be casted. 

## How To Run (development)

1. Install dependencies
```
go mod tidy
```

2. Run main file
```
go run main.go
```

3. Build main file
```
go build main.go
```


`NOTE (for maintainers):` The deployment and procedure notes are [here](https://github.com/thcollective/deployment-automation-tool/blob/main/PROCEDURE.md)

<p align="right">(<a href="#top">back to top</a>)</p>



