FROM golang:1.21.1

RUN go version
ENV GOPATH=/

# Expose our port to the world
##EXPOSE 50051
##EXPOSE 8888
#EXPOSE 8080

COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/main.go

CMD ["./app"]