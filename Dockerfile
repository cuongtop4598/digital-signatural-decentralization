FROM golang:1.16-buster as builder
ENV APP_USER app
ENV APP_HOME /app
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY . .
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o depocket

FROM alpine:3.7 as runing_image_builder
RUN apk --no-cache add curl

FROM runing_image_builder
ENV APP_USER app
ENV APP_HOME /app
RUN addgroup -S $APP_USER && adduser -S $APP_USER -G $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY --chown=$APP_USER:$APP_USER --from=builder $APP_HOME/depocket $APP_HOME