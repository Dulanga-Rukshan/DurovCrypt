#!/bin/bash


set -e

APP_NAME="DurovCrypt"
INSTALL_DIR="${HOME}/.local/bin"
BINARY_PATH="${INSTALL_DIR}/${APP_NAME}"

#colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

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
                    echo "Please install Go manually from https://golang.org/dl/"
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

#continue with installation tool
echo -e "${GREEN}=== Installing ${APP_NAME} ===${NC}"

#create a dir
mkdir -p "$INSTALL_DIR"

#build the durovcrypt
echo "Building binary..."
if ! go build -o "$APP_NAME"; then
    echo -e "${RED}Build failed${NC}"
    exit 1
fi

#insall
echo "Installing to ${BINARY_PATH}"
mv -f "$APP_NAME" "$BINARY_PATH"
chmod +x "$BINARY_PATH"

#Verfiy
if [[ -f "$BINARY_PATH" ]]; then
    echo -e "${GREEN}Successfully installed!${NC}"
else
    echo -e "${RED}Installation failed${NC}"
    exit 1
fi

#checking the path variable
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    echo -e "${YELLOW}Note: ${INSTALL_DIR} is not in your PATH${NC}"
    echo "To make the tool available everywhere, add this to your shell configuration:(~/.bashrc or ~/.zshrc file )"
    echo "export PATH=\"${INSTALL_DIR}:\$PATH\""
    echo "Then run: source ~/.bashrc (or your shell config file)"
fi

echo -e "${GREEN}Try running: ${APP_NAME}${NC}"