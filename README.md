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

* [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_darwin_amd64.tar.gz)
```
curl -vLJO https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_darwin_amd64.tar.gz
```
```
tar -xzvf https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_darwin_amd64.tar.gz
``` 
```
chmod +x ./thc-cli-tool
```
```
mv ./thc-cli-tool /usr/local/bin
```
* [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_darwin_arm64.tar.gz)
```
curl -vLJO https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_darwin_arm64.tar.gz
```
```
tar -xzvf https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_darwin_arm64.tar.gz 
```
```
chmod +x ./thc-cli-tool
```
```
mv ./thc-cli-tool /usr/local/bin
```

### Linux

* [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_linux_amd64.tar.gz)
```
curl -vLJO https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_linux_amd64.tar.gz
```
```
tar -xzvf https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_linux_amd64.tar.gz 
```
```
chmod +x ./thc-cli-tool
```
```
mv ./thc-cli-tool /usr/local/bin
```
* [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_linux_arm64.tar.gz)
```
curl -vLJO https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_linux_arm64.tar.gz
```
```
tar -xzvf https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_linux_arm64.tar.gz 
```
```
chmod +x ./thc-cli-tool
```
```
mv ./thc-cli-tool /usr/local/bin
```

### Windows
*  Tap here to download for [amd64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_windows_amd64.tar.gz)
*  Tap here to download for [arm64](https://github.com/thcollective/thc-deployment-tool/releases/download/v0.3.8/thc-deployment-tool_0.3.8_windows_arm64.tar.gz)
*  Right click the file and extract preferably using `7zip`.
*  Add `thc-cli-tool.exe` to system properties `Environment Variables...`

2. For Mac/Linux users, run `thc-cli-tool` using the `terminal` on root directory of your project, then start answering the prompt.

3. For windows user, run `./thc-cli-tool.exe` on your working project root directory, preferably run it on your `terminal`, then start answering the prompt.

4. The live url will be display on github actions `tab` located under `deploy` job.

5. Hit ctrl+c or ctrl+d to if you want to abort the process.

6. For redpanda users, note that `appName` is a global name and developers need to keep it consistent followed by application that they using to `pub/sub` a `message/event`

* create a topic from POST https://thc-redpanda-admin-d7mdbpojaa-as.a.run.app/api/v1/redpanda/createTopic

payload: 
```
{
	"clientId": "appName",
	"topic": "topicName",
	"numPartitions": 3
}
```

* To check what topics is available on the brokers, GET https://thc-redpanda-admin-d7mdbpojaa-as.a.run.app/api/v1/redpanda/listTopic to see what topics are available on the broker

* To delete a topic, DELETE https://thc-redpanda-admin-d7mdbpojaa-as.a.run.app/api/v1/redpanda/deleteTopic

payload:  
```
{
	"clientId": "appName",
	"topic": "topicName"
}
```

* Continue with thc-cli-tool prompt  

* Install dependencies @ `npm i kafkajs uuid dotenv @kafkajs/confluent-schema-registry`

* Get Redpanda Kafka Broker IP addresses and Redpanda Schema Registry Host URL from Adri or Ming

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



