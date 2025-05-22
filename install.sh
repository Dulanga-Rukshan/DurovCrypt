#!/bin/bash

set -e

app_name="DurovCrypt"
install_dir="${HOME}/.local/bin"
binary_path="${install_dir}/${app_name}"

# colors
red='\033[0;31m'
green='\033[0;32m'
yellow='\033[1;33m'
blue='\033[0;34m'
nc='\033[0m'

command_exists() {
    command -v "$1" >/dev/null 2>&1
}
#install go in linux if there no go
install_go_linux() {
    echo -e "${blue}installing go...${nc}"
    
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
        echo -e "${red}could not detect package manager.${nc}"
        echo "please install go manually from https://golang.org/dl/"
        exit 1
    fi
    
    if ! command_exists go; then
        echo -e "${red}installation failed. please install go manually.${nc}"
        exit 1
    fi
    
    echo -e "${green}go installed successfully!${nc}"
}
##install go in mac 
install_go_macos() {
    echo -e "${blue}installing go on macos...${nc}"
    
    if command_exists brew; then
        brew install go
    else
        echo -e "${yellow}homebrew not found.${nc}"
        echo "1) install homebrew first (recommended): https://brew.sh"
        echo "2) or download go directly from: https://golang.org/dl/"
        exit 1
    fi
    
    if ! command_exists go; then
        echo -e "${red}installation failed. please install go manually.${nc}"
        exit 1
    fi
    
    echo -e "${green}go installed successfully!${nc}"
}
#if go not exists ask to install it from user
if ! command_exists go; then
    echo -e "${yellow}go is not installed on your system.${nc}"
    echo -e "go is required to build and install ${app_name}."
    
    os=$(uname -s)
    
    while true; do
        read -rp "do you want to install go now? (y/n): " choice
        case "$choice" in
            y|Y )
                echo -e "${blue}attempting to install go...${nc}"
                
                if [ "$os" = "Linux" ]; then
                    install_go_linux
                elif [ "$os" = "Darwin" ]; then
                    install_go_macos
                else
                    echo -e "${yellow}automatic installation not supported for your os.${nc}"
                    echo "please install go manually from https://golang.org/dl/"
                    exit 1
                fi
                break
                ;;
            n|N )
                echo -e "${red}aborting installation. please install go manually and try again.${nc}"
                echo "download go from: https://golang.org/dl/"
                exit 1
                ;;
            * )
                echo "please answer y or n."
                ;;
        esac
    done
fi

echo -e "${green}=== installing ${app_name} ===${nc}"

mkdir -p "$install_dir"

echo "building binary..."
if ! go build -buildvcs=false -o "$app_name"; then
    echo -e "${red}build failed${nc}"
    exit 1
fi

echo "installing to ${binary_path}"
mv -f "$app_name" "$binary_path"
chmod +x "$binary_path"

if [[ -f "$binary_path" ]]; then
    echo -e "${green}successfully installed!${nc}"
else
    echo -e "${red}installation failed${nc}"
    exit 1
fi

if [[ ":$PATH:" != *":${install_dir}:"* ]]; then
    echo -e "${yellow}note: ${install_dir} is not in your path${nc}"
    echo "to make the tool available everywhere, add this to your shell configuration:"
    echo "export PATH=\"${install_dir}:\$PATH\""
    echo "then run: source ~/.bashrc (or your shell config file)"
fi

echo -e "${green}try running: ${app_name}${nc}"