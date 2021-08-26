MAINTAINER mohak.kant1amg@gmail.com

# Create the build Stage to compile our code
# Inherit from the golang 1.17 container
FROM golang:1.17 AS builder

# Update the system
RUN apt update -y

# Install upx that we will later use to
# compress our binary
RUN apt install upx -y

# Set the working directory to be within the
# golang default path
WORKDIR /go/src/github.com/mohakamg/fastinversesquareroot

# Copy the required files
COPY . .

# Build the binary
# Specify hardware architecture here if required
RUN cd cmd/fisrserver && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" .

# Compress the binary
RUN upx --brute ./cmd/fisrserver/fisrserver

# We create the stage to run the executable
# Since Go Binaries bring all the system dependanices
# with it, there is no need for an OS
FROM scratch

COPY --from=builder /go/src/github.com/mohakamg/fastinversesquareroot/cmd/fisrserver/fisrserver .
ENTRYPOINT ["./fisrserver"]
