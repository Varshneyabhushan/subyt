# syntax=docker/dockerfile:1

FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /ytWorker

FROM gcr.io/distroless/base-debian10 As Execution

WORKDIR /

COPY --from=build /ytWorker /ytWorker

USER nonroot:nonroot

ENTRYPOINT ["/ytWorker"]

COPY env.json ./

CMD [ "/ytWorker" ]