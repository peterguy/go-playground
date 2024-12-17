FROM --platform=linux/amd64 ubuntu:latest

RUN apt-get update && \
    apt-get install -y wget software-properties-common build-essential gnome-keyring

# Install Go
ENV GO_VERSION 1.22.3
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

# Set PATH to include Go binaries
ENV PATH $PATH:/usr/local/go/bin

ENV CGO_ENABLED=1

# Define /src as a mount point
VOLUME /src
# Builds a Docker image for a secure storage service.
# The image includes the necessary dependencies and configuration
# to run the secure storage service.
