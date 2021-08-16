FROM golang:1.15.5-alpine

# NOTE: GCC is needed to run the tests as is:
# - https://github.com/golang/go/issues/42996
RUN apk --no-cache add build-base

# NOTE golang has a folder naming convention:
WORKDIR /src

init:
  RUN go mod init summon-1password
	RUN go get github.com/1Password/connect-sdk-go
  SAVE ARTIFACT go.mod AS LOCAL go.mod
  SAVE ARTIFACT go.sum AS LOCAL go.sum

deps:
  COPY . .
  RUN go mod tidy
  SAVE ARTIFACT go.mod AS LOCAL go.mod
  SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
  FROM +deps
	COPY cmd .
  RUN go build -o summon-1password
  # NOTE BUILD CONTAINER

image:
  COPY +build/summon-1password .
  ENTRYPOINT ["/summon-1password"]
  SAVE IMAGE summon-1password:latest

test:
  FROM +deps
  RUN go test -v -coverprofile=app.coverage
  RUN go tool cover -func=app.coverage -html
  SAVE ARTIFACT coverage.html AS LOCAL coverage.html
