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
* Cloud Run github actions with .env

## How To Run 

1. Download the latest binary package [here](https://github.com/thcollective/thc-deployment-tool/releases) depending on your machine (Mac, Linux, Windows)

2. Extract the folder and drag the file named `cli` to your root directory of your working project.

3. Run `./cli` on your working project root directory, preferably run it on your `cli`, then start answering the prompt.

`NOTE (for maintainers):` The procedure and deployment notes are [here](https://github.com/thcollective/thc-deployment-tool/blob/main/PROCEDURE.md)


### How To Run (in development mode)

1. Install dependencies
```
go mod tidy
```

2. Run main file
```
go run cmd/cli/main.go
```

3. Build main file
```
go build cmd/cli/main.go
```



<p align="right">(<a href="#top">back to top</a>)</p>



