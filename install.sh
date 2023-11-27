#!/bin/bash

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'
blue='\033[0;34m'

# Check if Python is installed
if ! command -v python3 >/dev/null 2>&1; then
  echo ${red}"Python3 is not installed. Please install Python3 and try again."
  exit 1
fi

# Check if Docker is installed
if [ -x "$(command -v docker)" ]; then
  echo ${red}"Docker is not installed. Please install Docker and try again."
  exit 1
fi

# Stop old image and rebuild
echo ${blue}"Stop running last 3x-ui-bot image..."
$(docker stop $(docker ps -a -q --filter ancestor=3x-ui-bot --format="{{.ID}}"))

#echo ${blue}"Installing 3x-ui-bot docker image..."
docker compose up --build -d 3x-ui-bot

# Install Python dependencies using pip
echo ${blue}"Installing Python dependencies..."
pip3 install -r requirements.txt

# Display Python and dependency versions
python_version=$(python3 --version 2>&1)
dependency_version=$(your_dependency --version 2>&1)

# Replace installed directory
this_path=$(pwd)
command=$(sed "\|installed_directory=| s|.*|installed_directory=$this_path|" 3x-ui-bot.sh)
echo "$command" > 3x-ui-bot.sh

# Install 3x-ui-bot
echo ${blue}"Installing 3x-ui bot..."
chmod +x 3x-ui-bot.sh
sudo cp 3x-ui-bot.sh /usr/local/bin/3x-ui-bot

echo ${green}"Python version: $python_version"
echo ${green}"Installation and version check completed successfully."
