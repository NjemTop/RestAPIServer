export APP_NAME=rest-api-server

run :
    go run .

build-docker:
    docker build -t ${APP_NAME} .

run-docker:
    docker run -d -p 8055:8055 ${APP_NAME}

run-comsope:
    docker-comsope up -d

down-comsope:
    docker-comsope down
