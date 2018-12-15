FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD temp/dep /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/HighLoad18
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
ENV GIN_MODE=release
ENV SERVER_PORT=":80"
ENV ZIP_FILE=/tmp/data/data.zip
COPY --from=builder /server ./server
COPY ./temp/data/data.zip /tmp/data/data.zip
EXPOSE 80
ENTRYPOINT ["./server"]