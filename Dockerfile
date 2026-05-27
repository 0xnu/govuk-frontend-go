FROM golang:1.22-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/example ./cmd/example-gin-app

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=build /out/example /example
EXPOSE 91942
USER nonroot:nonroot
ENTRYPOINT ["/example"]

