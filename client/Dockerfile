FROM golang:alpine as build-env
RUN apk --no-cache add build-base git bzr mercurial gcc portaudio portaudio-dev pkgconfig
ADD . /src
WORKDIR /src
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -o /src/client

FROM alpine
RUN apk --no-cache add portaudio
COPY --from=build-env /src/client /client
ENTRYPOINT ["/client"]