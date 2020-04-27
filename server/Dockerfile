FROM golang:alpine as build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
WORKDIR /src
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/server

FROM alpine
COPY --from=build-env /src/server /server
ENTRYPOINT ["/server"]