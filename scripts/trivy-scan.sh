#!/bin/bash

install_trivy() {
    sudo apt-get install wget apt-transport-https gnupg lsb-release -y
    wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
    echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee -a /etc/apt/sources.list.d/trivy.list
    sudo apt-get update
    sudo apt-get install trivy -y
}

execute_trivy() {
    trivy image --input ${IMAGE_PATH}
}

main() {
    install_trivy
    execute_trivy
}

main
