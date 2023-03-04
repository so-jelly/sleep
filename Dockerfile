FROM golang:1.18 as builder
WORKDIR /workspace
COPY main.go go.mod /workspace/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o sleep main.go

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/sleep .
USER 65532:65532

ENTRYPOINT ["/sleep"]
