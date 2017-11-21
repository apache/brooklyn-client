
# [![**Brooklyn**](https://brooklyn.apache.org/style/img/apache-brooklyn-logo-244px-wide.png)](http://brooklyn.apache.org/)

### Apache Brooklyn Client Sub-Project

This repo contains the CLI client for Apache Brooklyn.
It is written in go and is built for all platform.

### Building the project

2 methods are available to build this project: within a docker container or directly with maven.

#### Using docker

The project comes with a `Dockerfile` that contains everything you need to build this project.
First, build the docker image:

```bash
docker build -t brooklyn:client .
```

Then run the build:

```bash
docker run -i --rm --name brooklyn-client -v ${HOME}/.m2:/root/.m2 -v ${PWD}:/usr/build -w /usr/build brooklyn:client mvn clean install
```

### Using maven

Simply run:

```bash
mvn clean install
```