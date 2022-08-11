FROM golang:1.17-alpine as build-env
 
ENV APP_NAME request-inspector
ENV CMD_PATH main.go
 
COPY . $GOPATH/src/$APP_NAME
COPY ./templates /templates
WORKDIR $GOPATH/src/$APP_NAME
 
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
FROM alpine:3.14
 
ENV APP_NAME request-inspector
 
COPY --from=build-env /$APP_NAME  .
COPY --from=build-env /templates  ./templates
 
EXPOSE 8080

CMD ./$APP_NAME
