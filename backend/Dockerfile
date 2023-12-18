FROM amd64/golang:1.21.5-bookworm AS build

WORKDIR /go/src/backend

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV ARCH=amd64

COPY . .
RUN go mod download

RUN go build -o ./bin/backend ./cmd/app/main.go

FROM debian:bookworm AS runtime
WORKDIR / 

COPY --from=build /go/src/backend/bin/backend /go/bin/

EXPOSE 8000
ENTRYPOINT ["/go/bin/backend", "-db=/config/umo.db", "-jwt=/config/jwt.key"]
