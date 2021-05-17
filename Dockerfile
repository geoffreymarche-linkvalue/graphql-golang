FROM golang:1.16 as builder

ENV BP=$GOPATH/src/graphql-golang
WORKDIR $BP
COPY go.mod .
COPY go.sum .
ENV GOPROXY=direct
ENV GOSUMDB=off
RUN go mod download
COPY ./internal ./internal
COPY ./mutations ./mutations
COPY ./queries ./queries
COPY ./security ./security
COPY ./types ./types
COPY ./main.go ./main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /bin/app

# Build final image
FROM scratch
WORKDIR /bin/
COPY --from=builder /bin/app .
COPY ./config.json .
EXPOSE 4000

CMD ["/bin/app", "run"]