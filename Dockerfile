# Builder
FROM golang:1.15.6 as build

ENV TARGET_OS linux
ENV TARGET_ARCH amd64

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -o main

# Runner
FROM ubuntu:20.04

RUN apt-get update && apt-get install -y ca-certificates --no-install-recommends && rm -rf /var/lib/apt/lists/*

ENV APP_USER user
ENV APP_HOME /app

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p /usr/src/app
WORKDIR ${APP_HOME}

COPY --from=build /app/main ${APP_HOME}/main

USER ${APP_USER}
CMD ["./main"]