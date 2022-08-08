# boilerplate

## Requirements
* Golang 1.17 or newer
* Git
* NMP - if you want to generate api 
```npm install @openapitools/openapi-generator-cli -g```

### Environment Variables
```
ENV = local, prod, or dev
``` 

### Build Project
```
make build
```  

### Run Project
```
make run
```

### Package Project for AWS to Deploy With CodePipeline
```
make package-auto
```  

### Package Project for AWS to Deploy Manually
```
make package-manual
```  

### Run Integration Tests
```
make test-integration
```

### Run Unit Tests
```
make test-unit
```

### Generate API Docs
```
make api-gen
```

### Setting Up CloudWatchLogs (for log metrics)
```

Make sure the Elastic Beanstalk's IAM profile has a role that gives it full permission for CloudWatchLogs

Currently the .ebextensions has a config file that will automatically setup CloudWatchLogs for you. However if yuou need
to set it up manually do the following:

Open up the configuration file:
sudo vim /etc/awslogs/config/beanstalklogs.conf

Add the following to the file:

[/var/log/web-1.log]
log_group_name=`{"Fn::Join":["/", ["/aws/elasticbeanstalk", { "Ref":"AWSEBEnvironmentName" }, "var/log/web-1.log"]]}`
log_stream_name={instance_id}
file=/var/log/web-1.log*
[/var/log/web-1.error.log]
log_group_name=`{"Fn::Join":["/", ["/aws/elasticbeanstalk", { "Ref":"AWSEBEnvironmentName" }, "var/log/web-1.error.log"]]}`
log_stream_name={instance_id}
file=/var/log/web-1.error.log*

Since the aws log service is probably already running restart it:
sudo service awslogs restart
```


### Setting Up CodePipeline (build and,deploy)
Step 1) Create A Pipeline with the same fields as below

![Step1](/readme_resources/Step1.png)

Step 2)  Next, select Github as the source provider. You will need to connect your Github account to select a repo. 

![Step2](/readme_resources/Step2.png)

Step 3) Make sure to select webhooks in the detection options

![Step3](/readme_resources/Step3.png)

Step 4) In the build stage select AWSCodeBuild as the build provider. You will need a CodeBuild project so click on the "Create Project" button

![Step4](/readme_resources/Step4.png)

Step 5) Fill in the options same as the following. 

![Step5A](/readme_resources/Step5A.png)

If your project uses Environment variables add them like in the following image 

![Step5B](/readme_resources/Step5B.png)
![Step5C](/readme_resources/Step5C.png)

After this you will be back to finishing up the build state in CodePipeline

![Step6](/readme_resources/Step6.png)

Step 6) In the deploy stage, make sure you have an Elasticbeanstalk application setup and select it 

![Step6](/readme_resources/Step7.png)

### Adding Integration Tests to BuildPipeline

1. On the CodePipeline click edit
2. Add a stage between Source and Build. Call This stage Test
3. Add an action called "Integration Tests"
4. For Action Provider set "AWS Codebuild"
5. For Input Artifact set "SourceArtifact"
6. Create a CodeBuild project that is the same as Step 5 in previous instructions. Only difference is buildfile will be "buildsepcs/integration-tests.yml"

You will end up with a pipeline similar to this
![Pipeline Overview](/readme_resources/PipelineOverview.png)

### Adding Unit Tests to BuildPipeline
1. Repeat steps in previous instructions but this time buildfile will be "buildsepcs/unit-tests.yml"



### Setting up cognito
![Cognito1](/readme_resources/cognito1.png)
![Cognito1](/readme_resources/cognito2.png)
![Cognito1](/readme_resources/cognito3.png)

Make sure that when you are creating the roles for cognito identities that you give the broadest permissions that you will need. On the backend when you attach a policy the attached policy will restrict what devices can be accessed
