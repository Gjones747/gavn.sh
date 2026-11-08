FROM golang:1.27 AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /out/server .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/server /server
USER nonroot
ENTRYPOINT ["/server"]
