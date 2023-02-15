FROM golang:1.19-alpine as build

WORKDIR /source
COPY . .

ARG COMMIT
RUN #CGO_ENABLED=0 go build -ldflags "-s -w -X main.commit=${COMMIT}" -o bin/pipeline main.go
RUN CGO_ENABLED=0 go build -o bin/main main.go

FROM alpine:3.12

COPY --from=build /source/bin/main /bin/main

EXPOSE 8081

ENTRYPOINT ["./bin/main"]
