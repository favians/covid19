FROM golang:alpine
LABEL maintainer="golang_starter"

ENV APP_PATH = "/app"

WORKDIR ${APP_PATH}

COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8000

CMD [ "go", "run", "main.go", "run" ]