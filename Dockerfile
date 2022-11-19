# --- Stage 1:
FROM golang:1.19-alpine as builder
# Args & ENVs
ENV BUILD_PATH=/go/src/github.com/NjemTop/automatic_email

# COPY local files
WORKDIR /app
COPY . .

# Get go dependencies
RUN go mod download

RUN go run .

EXPOSE 8055