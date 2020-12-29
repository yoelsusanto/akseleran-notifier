FROM golang:1.15.6-alpine as build

ENV APP_USER user
ENV APP_HOME /go/src/akseleran-notifier/
ENV TARGET_OS linux
ENV TARGET_ARCH amd64

# RUN groupadd ${APP_USER} && useradd -m -g ${APP_USER} -l ${APP_USER}
RUN mkdir -p ${APP_HOME}

WORKDIR ${APP_HOME}
# USER ${APP_USER}
COPY . .

RUN go mod download
RUN go mod verify
RUN GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -o akseleran-notifier

FROM ubuntu:20.04

ENV APP_USER user
ENV APP_HOME /go/src/akseleran-notifier/

RUN groupadd ${APP_USER} && useradd -m -g ${APP_USER} -l ${APP_USER}
RUN mkdir -p ${APP_HOME}

COPY --from=build ${APP_HOME}/akseleran-notifier ${APP_HOME}

USER ${APP_USER}
CMD ["./akseleran-notifier"]