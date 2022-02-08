<div align="center">
  <a href="https://github.com/thcollective">
    <img src="img/thc.png" alt="thc_logo" width="80" height="80">
  </a>
  <h3 align="center">THC deployment automation tool</h3>
  <h3 align="center">Maintained by: <a href="https://github.com/ass77">ass77</a></h3>
</div>

## What is it about?

This deployment automation tool helps users to generate basic files that are required to deploy their app on cloud run, scan their code to check for bugs before deployment and tracing any error on the app in dev mode or prod mode. The script will generate files as follows:

* Dockerfile for frontend (vue, nuxt) and backend(expressJS, gofiber)
* Sonarcloud github actions and properties file (need to seek from Admin for keys)
* Cloud Run github actions with environments
* Pub/Sub to Redpanda Kafka 

## Download and Run 

1. Download the binary package based on your OS and set to local path.

### Darwin
*  [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_darwin_amd64.tar.gz)
```
curl -Lo ./thc-toolkit https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_darwin_amd64.tar.gz && \
mv ./thc-toolkit /usr/local/bin
``` 
*  [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_darwin_arm64.tar.gz)
```
curl -Lo ./thc-toolkit https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_darwin_arm64.tar.gz && \
mv ./thc-toolkit /usr/local/bin
```

### Linux
*  [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_linux_amd64.tar.gz)
```
curl -Lo ./thc-toolkit https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_linux_amd64.tar.gz && \
mv ./thc-toolkit /usr/local/bin
```
*  [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_linux_arm64.tar.gz)
```
curl -Lo ./thc-toolkit https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_linux_arm64.tar.gz && \
mv ./thc-toolkit /usr/local/bin
```

### Windows
*  [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_windows_amd64.tar.gz)
*  [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.6/thc-deployment-tool_0.3.6_windows_arm64.tar.gz)

2. Extract the binary and save to local path (for Windows)

* Right click the file and extract preferably using `7zip`.
* Add `thc-cli-tool.exe` to system properties `Environment Variables...`

3. For Mac/Linux users, run `thc-toolkit` using `terminal` on root directory of your project, then start answering the prompt.

4. For windows user, run `./thc-cli-tool.exe` on your working project root directory, preferably run it on your `terminal`, then start answering the prompt.

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
go run cmd/thc-cli-tool/main.go
```

3. Build main file
```
go build cmd/thc-cli-tool/main.go
```



<p align="right">(<a href="#top">back to top</a>)</p>



