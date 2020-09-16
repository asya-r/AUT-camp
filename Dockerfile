FROM golang:1.13.5
WORKDIR /app
ADD . /app
ENV GOPATH /app

CMD ["go", "run", "/app/src/main.go"]
