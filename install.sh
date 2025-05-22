#!/bin/bash

set -e

app_name="DurovCrypt"


# colors
red='\033[0;31m'
green='\033[0;32m'
yellow='\033[1;33m'
blue='\033[0;34m'
nc='\033[0m'

#detect user
if [ "$(id -u)" -eq 0 ]; then
    install_dir="/usr/local/bin" #root
else
    install_dir="${HOME}/.local/bin"  #normal users
fi

binary_path="${install_dir}/${app_name}"

command_exists() {
    command -v "$1" >/dev/null 2>&1
}

add_to_path() {
    
    [ "$(id -u)" -eq 0 ] && return
    
    local shell_rc
    #find the right shell config
    [ -f "${HOME}/.zshrc" ] && shell_rc="${HOME}/.zshrc" || shell_rc="${HOME}/.bashrc"
    
    #create if doesnt exist
    [ -f "$shell_rc" ] || touch "$shell_rc"
    
    #add to path if not already present
    if ! grep -q "export PATH=\"${install_dir}:\$PATH\"" "$shell_rc"; then
        echo "export PATH=\"${install_dir}:\$PATH\"" >> "$shell_rc"
        echo -e "${green}→ Added ${install_dir} to PATH in ${shell_rc}${nc}"
    fi
    
    #update the current shell
    export PATH="${install_dir}:$PATH"
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
        case "$choice" in y|Y )
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
    echo -e "${red}DurovCrypt build failed${nc}"
    exit 1
fi


if [ ! -f "$app_name" ]; then
    echo -e "${red}built binary not found${nc}"
    exit 1
fi

echo "installing to ${binary_path}"
mv -f "$app_name" "$binary_path"
chmod +x "$binary_path"


if [[ -f "$binary_path" ]]; then
    echo -e "${green}DurovCrypt is successfully installed!${nc}"
    
    if [[ ":$PATH:" != *":${install_dir}:"* ]]; then
        echo -e "${yellow}adding ${install_dir} to your path...${nc}"
        add_to_path
    fi

    echo -e "${green}try running: ${app_name}${nc}"
else
    echo -e "${red}installation failed${nc}"
    exit 1
fi