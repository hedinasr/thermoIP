To build docker image:
`docker build -t mygoapp .`

To run:
`docker run --rm -it -p 8080:8080 -v \
$(pwd)/myGoDockerApp/app/:/go/src/github.com/jdoe/myProject/app \
mygoapp`

The app is `app/tmp/runner-build`
