# syntax=docker/dockerfile:1

FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /ytService

FROM gcr.io/distroless/base-debian10 As Execution

WORKDIR /

COPY --from=build /ytService /ytService

USER nonroot:nonroot

ENTRYPOINT ["/ytService"]

COPY env.json ./

CMD [ "/ytService" ]