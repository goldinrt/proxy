FROM golang:1.17-alpine
WORKDIR /usr/src/app
COPY . .
RUN go build
RUN rm -f *.go
EXPOSE 8080
ENTRYPOINT [ "./proxy" ] 
