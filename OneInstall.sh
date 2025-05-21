#!/bin/bash

# Enhanced install script for my-go-tool

set -e

# Configuration
APP_NAME="DurovCrypt"
VERSION="1.0.0"
INSTALL_DIR="/usr/local/bin"
REPO_URL="https://github.com/Dulanga-Rukshan/DurovCrypt.git"
BUILD_DIR=$(mktemp -d)

#colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'


error() {
  echo -e "${RED}[ERROR]${NC} $1" >&2
  exit 1
}

info() {
  echo -e "${GREEN}[INFO]${NC} $1"
}

warn() {
  echo -e "${YELLOW}[WARN]${NC} $1"
}

#check if the commands exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

#install go 
install_go_linux() {
    echo -e "${BLUE}Installing Go...${NC}"
    
    #check package manager 
    if command_exists apt; then
        sudo apt update
        sudo apt install -y golang
    elif command_exists yum; then
        sudo yum install -y golang
    elif command_exists dnf; then
        sudo dnf install -y golang
    elif command_exists zypper; then
        sudo zypper install -y go
    elif command_exists pacman; then
        sudo pacman -Sy --noconfirm go
    else
        echo -e "${RED}Could not detect package manager.${NC}"
        echo "Please install Go manually from https://golang.org/dl/"
        exit 1
    fi
    
    #verifiyng the installation
    if ! command_exists go; then
        echo -e "${RED}Installation failed. please install Go manually.${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}Go installed successfully!${NC}"
}


#root privilege
if [ "$(id -u)" -ne 0 ]; then
  warn "This script require root privileges for installation"
fi



# Check dependencies
info "Checking dependencies..."
check_dependency() {
  if ! command -v "$1" >/dev/null 2>&1; then
    error "'$1' is required but not installed. Please install it first."
  fi
}

check_dependency git
check_dependency go

# Clone repository
info "Cloning repository..."
git clone --quiet --branch "v$VERSION" --depth 1 "$REPO_URL" "$BUILD_DIR" || {
  error "Failed to clone repository"
}



#install go on mac os 
install_go_macos() {
    echo -e "${BLUE}Installing Go on macOS...${NC}"
    
    if command_exists brew; then
        brew install go
    else
        echo -e "${YELLOW}Homebrew not found.${NC}"
        echo "1) Install Homebrew first (recommended): https://brew.sh"
        echo "2) Or download Go directly from: https://golang.org/dl/"
        exit 1
    fi
    
   #verifiyng the installation
    if ! command_exists go; then
        echo -e "${RED}Installation failed. Please install Go manually.${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}Go installed successfully!${NC}"
}

#check for go 
if ! command_exists go; then
    echo -e "${YELLOW}Go is not installed on your system.${NC}"
    echo -e "Go is required to build and install ${APP_NAME}."
    
    #os 
    OS=$(uname -s)
    
    while true; do
        read -rp "Do you want to install Go now? (y/n): " choice
        case "$choice" in
            y|Y )
                echo -e "${BLUE}Attempting to install Go...${NC}"
                
                if [ "$OS" = "Linux" ]; then
                    install_go_linux
                elif [ "$OS" = "Darwin" ]; then
                    install_go_macos
                else
                    echo -e "${YELLOW}automatic installation not supported for your OS.${NC}"
                    echo "install Go manually from https://golang.org/dl/"
                    exit 1
                fi
                break
                ;;
            n|N )
                echo -e "${RED}Aborting installation. please install Go manually and try again.${NC}"
                echo "Download Go from: https://golang.org/dl/"
                exit 1
                ;;
            * )
                echo "Please answer y or n."
                ;;
        esac
    done
fi



cd "$BUILD_DIR" || error "failed to enter build directory"

#build the tool
info "Building $APP_NAME..."
go build -ldflags="-s -w" -trimpath -o "$APP_NAME" . || {
  error "Build failed"
}


info "Verifying the binary..."
./"$APP_NAME" --version || {
  error "Built binary verification failed"
}

#install the tool
info "Installing $APP_NAME to $INSTALL_DIR..."
sudo mkdir -p "$INSTALL_DIR"
sudo mv -f "$APP_NAME" "$INSTALL_DIR" || {
  error "Installation failed - try running with sudo"
}

#set permissions to tool
sudo chmod 755 "$INSTALL_DIR/$APP_NAME"


info "Cleaning up..."
rm -rf "$BUILD_DIR"

#verify the installation
if command -v "$APP_NAME" >/dev/null 2>&1; then
  info "Installation complete! You can now run '$APP_NAME' from anywhere."
  info "Try: $APP_NAME --help"
else
  warn "Installation seems successful but the binary isn't in your PATH"
  warn "Please ensure $INSTALL_DIR is in your PATH environment variable"
  warn "To make the tool available everywhere, add this to your shell configuration:(~/.bashrc or ~/.zshrc file )"
  warn "export PATH=\"${INSTALL_DIR}:\$PATH\""
  warn "Then run: source ~/.bashrc (or your shell config file)"
fi