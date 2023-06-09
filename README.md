# 3x-ui bot

This bot tells clients traffic usage by config url.

It is working for `vless` and `vmess` configs now.

## how to run:
1. install docker:
  - `sudo apt-get update`
  - `sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin`
2. get project:
   - `git clone https://github.com/Bright98/3x-ui-bot.git`
   - `cd 3x-ui-bot`
   - change `.env` file:
     - bot token
     - server ip
     - panel port
     - panel username
     - panel password

   - `sh install.sh`
