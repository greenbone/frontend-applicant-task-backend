FROM golang:1.20 AS BUILDER
MAINTAINER Mika Greif <mika.greif@greenbone.net>
WORKDIR /app
COPY ./go.mod ./go.sum ./
COPY ./src ./src

RUN go version
RUN go mod download
RUN CGO_ENABLED=0 go build -a -o ./server ./src/main.go

FROM alpine:3.18

WORKDIR /app
COPY --from=BUILDER /app/server ./server

RUN addgroup -S appuser && adduser -S appuser -G appuser
RUN ls -laR

RUN chown appuser:appuser ./server
USER appuser
EXPOSE 8080
RUN ls -laR

CMD ["/app/server"]
