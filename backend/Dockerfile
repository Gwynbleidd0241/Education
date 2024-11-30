FROM golang:1.23-alpine AS build

WORKDIR /EduCourse

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /EduCourse/app ./cmd/EduCourse

FROM alpine:latest
WORKDIR /app

COPY --from=build /EduCourse/app /EduCourse/app
COPY ./logs ./logs
COPY ./config/config.yaml ./config/config.yaml

EXPOSE 8080
CMD ["/EduCourse/app"]
