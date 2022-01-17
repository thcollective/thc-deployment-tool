<div align="center">
  <a href="https://github.com/thcollective">
    <img src="img/thc.png" alt="thc_logo" width="80" height="80">
  </a>
  <h3 align="center">THC deployment automation tool</h3>
</div>

## What is it about?

This deployment automation tool helps users to generate basic files that are required to deploy their app on cloud run, scan their code to check for bugs before deployment and tracing any error on the app in dev mode or prod mode. The script will generate files as follows:

* Dockerfile for frontend (vue, nuxt) and backend(expressJS, gofiber)
* Sonarcloud github actions and properties file (need to seek from Admin for keys)
* Cloud Run github actions with .env

## Download and Run 

1. Download the latest binary package depending on your machine (Mac, Linux, Windows)

### Darwin
*  [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.1/thc-deployment-tool_0.3.1_darwin_amd64.tar.gz) 
*  [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.0/thc-deployment-tool_0.3.0_darwin_arm64.tar.gz)

### Linux
*  [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.0/thc-deployment-tool_0.3.0_linux_amd64.tar.gz)
*  [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.0/thc-deployment-tool_0.3.0_linux_arm64.tar.gz)

### Windows
*  [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.0/thc-deployment-tool_0.3.0_windows_amd64.tar.gz)
*  [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.0/thc-deployment-tool_0.3.0_windows_arm64.tar.gz)

2. Extract the binary

### Mac
* `tar -xvzf thc-deployment-tool_0.3.0_linux_amd64.tar.gz -C <your_root_project_directory>`

### Linux
* `tar -xvzf thc-deployment-tool_0.3.0_linux_amd64.tar.gz -C <your_root_project_directory>`

### Windows
* Right click the file and extract it the folder to your root directory of your working project.
* or use `tar -xvzf x.tar.gz -C <your_root_project_directory>` if you have git bash installed on your windows cli


3. Remove `README.md` from the `.tar.gz` to avoid any conflicts with your project `README.md`. 

4. Run `./cli` on your working project root directory, preferably run it on your `cli`, then start answering the prompt.

5. The live url will be display on github actions `tab` located under `deploy` job.

6. Hit ctrl+c or ctrl+d to if you want to abort the process.

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



