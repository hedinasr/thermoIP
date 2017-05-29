FROM golang

COPY ./app /go/src/github.com/ananasr/myProject/app
WORKDIR /go/src/github.com/ananasr/myProject/app

RUN go get ./
RUN go build

CMD go get github.com/pilu/fresh && fresh

EXPOSE 8080
