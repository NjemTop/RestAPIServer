FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /RestAPIServer

RUN go get github.com/gin-gonic/gin

EXPOSE 8055

CMD [ "/RestAPIServer" ]
