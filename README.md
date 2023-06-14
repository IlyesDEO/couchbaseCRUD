# Couchbase CRUD 
## Installation
### Apple Silicon 
1. Install Docker
    * Download Docker Desktop for Mac ARM64 from [Docker Hub](https://desktop.docker.com/mac/main/arm64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-arm64)
    * Double-click `Docker.dmg` to open the installer, then drag the Docker icon to the Applications folder.
    * Double-click Docker.app in the Applications folder to start Docker.
    * The Docker menu (whale menu) displays the Docker Subscription Service Agreement window. Select Accept to continue.
    * From the installation window, select: \
    Use recommended settings (Requires password). This let’s Docker Desktop automatically set the necessary configuration settings.

2. Install Brew
    * Open Terminal and run the following command and follow instructions: \
    `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
    * Verify the installation by running the following command: \
    `brew --version`
3. Install GoLang
    * Open Terminal and run the following command: \
    `brew install go`
    * Verify the installation by running the following command: \
    `go version`


## Usage
See [Postman doc](./postman.md) to test the API. \
You also can downloand the [Postman collection](./postman.json) and import it to your Postman.
