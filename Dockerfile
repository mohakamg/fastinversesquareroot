# Create the build Stage to compile our code
# Inherit from the golang 1.17 container
FROM golang:1.17 AS builder

# Update the system
RUN apt update -y

# Install Snap
RUN apt install snapd -y

# Install upx that we will later use to
# compress our binary
RUN snap install upx -y

# Set the working directory to be within the
# golang default path
WORKDIR /go/src/github.com/mohakamg/fastinversesquareroot

# Copy the required files
COPY . .

# Build the binary
# Specify hardware architecture here if required
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" .

# Compress the binary
RUN upx --brute ./main

# We create the stage to run the executable
# Since Go Binaries bring all the system dependanices
# with it, there is no need for an OS
FROM scratch

COPY --from=builder /go/src/github.com/mohakamg/fastinversesquareroot/main .
ENTRYPOINT ["./main"]