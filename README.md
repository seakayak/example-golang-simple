## Seakayak Go example

This repository contains a simple Go web app that can be deployed to [seakayak.io](https://seakayak.io/).

### Prerequisites
- A seakayak.io account ([sign up here](https://seakayak.io/signup))
- A working Golang installation (see [instructions](https://golang.org/dl/))
- A working Docker client (see instructions for [mac](https://docs.docker.com/engine/installation/mac/), [linux](https://docs.docker.com/engine/installation/linux/), [windows](https://docs.docker.com/engine/installation/windows/))

### Quickstart

In the following, replacing `USERNAME` with your seakayak username.

```shell
# log into seakayak (one time only)
docker login seakayak.io

# clone the repository
git clone github.com/seakayak/example-golang-uniqueips
cd golang-example

# must build for linux because it is going into a Docker container
GOOS=linux go build -o linux-main *.go

# build the docker image
docker build -t seakayak.io/USERNAME/example-golang-uniqueips . 

# deploy your app
docker push seakayak.io/USERNAME/example-golang-uniqueips
```

Your app is now live. To see it, go to https://example-golang-uniqueips.USERNAME.seakayak.io/.

Next: explore the [documentation](https://seakayak.io/docs/).
