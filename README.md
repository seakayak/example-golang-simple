### Seakayak Go example

This repository contains a simple Go web app that can be deployed to [seakayak.io](https://seakayak.io/).

To build and deploy the app you will first need to [sign up for a seakayak account](https://seakayak.io/signup). You will also need a working [Docker client](https://docs.docker.com/engine/installation/mac/). Then run the following, replacing "USERNAME" with your seakayak username.

```shell
git clone github.com/seakayak/golang-example
cd golang-example

# must build for linux because it is going into a Docker container
GOOS=linux go build -o linux-main *.go

# build the docker image
docker build -t seakayak.io/USERNAME/golang-example . 

# deploy your app
docker push seakayak.io/USERNAME/golang-example
```

Now open https://golang-exampe.USERNAME.seakayak.io/ in your web browser.
