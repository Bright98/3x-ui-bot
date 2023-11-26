#!/bin/bash

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

# Check if Python is installed
if ! command -v python3 >/dev/null 2>&1; then
  echo ${red}"Python3 is not installed. Please install Python3 and try again."
  exit 1
fi

# Check if Docker is installed
#if [ -x "$(command -v docker)" ]; then
#  echo ${red}"Docker is not installed. Please install Docker and try again."
#  exit 1
#fi

# Stop old image and rebuild
#echo ${yellow}"Stop running last 3x-ui-bot image..."
#$(docker stop $(docker ps -a -q --filter ancestor=3x-ui-bot --format="{{.ID}}"))

#echo ${yellow}"Installing 3x-ui-bot docker image..."
#docker compose up --build -d 3x-ui-bot

# Install Python dependencies using pip
echo ${yellow}"Installing Python dependencies..."
pip3 install -r requirements.txt

# Display Python and dependency versions
python_version=$(python3 --version 2>&1)
dependency_version=$(your_dependency --version 2>&1)

# Install 3x-ui-bot
if [[ -e /usr/local/x-ui/ ]]; then
    systemctl stop 3x-ui-bot
    sudo rm /usr/local/3x-ui-bot/ -rf
fi
sudo cp 3x-ui-bot.sh /usr/local/bin/

chmod +x 3x-ui-bot.sh




echo ${green}"Python version: $python_version"
echo ${green}"Installation and version check completed successfully."
