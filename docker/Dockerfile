# First stage: build the application
FROM golang:1.14-alpine AS builder

ARG GOARCH
ARG GOARM

# Install build utilities
RUN apk --no-cache add --virtual .build-deps \
    bash \
    make \
    git

# Don't use ca-certificates as .build-deps so that we use import certificates
# from builder, in the final image
RUN apk --no-cache add \
    ca-certificates \
    && update-ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/ramonhpr/knotc

# Copy the source code from the current directory to $WORKDIR (inside the container)
COPY . .

# Build the application
RUN ./build.sh

# Remove build dependencies
RUN apk del .build-deps

# Second stage: create the entrypoint to the application binary generated in the previous stage
FROM scratch

WORKDIR /root/

# Copy SSL CA certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary file from the build stage
COPY --from=builder /go/src/github.com/ramonhpr/knotc/knotc knotc

ENTRYPOINT ["./knotc"]