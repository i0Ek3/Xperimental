#!/bin/bash

install_docker() {
    # Uninstall old versions
    sudo apt-get remove docker docker-engine docker.io containerd runc

    # Set up the repository
    # Update the apt package index and install packages to allow apt to use a repository over HTTPS
    sudo apt-get update
    sudo apt-get install \
        ca-certificates \
        curl \
        gnupg \
        lsb-release

    # Add Docker’s official GPG key
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

    # Use the following command to set up the repository
    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    # Update the apt package index, and install the latest version of Docker Engine, containerd, and Docker Compose, or go to the next step to install a specific version
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
}

build_and_run() {
    docker build -t linux-dev:latest .
    docker container run --rm -it linux-dev:latest /bin/sh
    #docker compose up -d
}

main() {
    platform=$(uname -s)
    if [ $platform == "Linux" ]
    then
        install_docker
    fi
    build_and_run
}

main
