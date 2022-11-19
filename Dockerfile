FROM golang:latest

COPY ./ ./

EXPOSE 8055

CMD [ "go", "run main.go"]
