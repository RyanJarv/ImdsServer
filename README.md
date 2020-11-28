# ImdsServer

This is a repo that deploys [RyanJarv/amazon-ec2-metadata-mock](https://github.com/RyanJarv/amazon-ec2-metadata-mock) to a lambda function/API gateway.

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process

### Installing dependencies & building the target 

```shell
make
```

### Local development


```bash
make invoke
```

## Packaging and deployment

To deploy while changing/checking any of the default settings you can first run:

```bash
sam deploy --guided
```

To deploy after that you can run:

```bash
make deploy
```

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
go test -v ./hello-world/
```
