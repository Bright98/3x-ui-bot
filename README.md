# 3x-ui bot

This bot tells clients traffic usage by config url.

It is working for `vless` and `vmess` configs now.

## How to run:
1. install docker:
  - `bash <(curl -sSL https://get.docker.com)`
2. get project:
   - `git clone https://github.com/Bright98/3x-ui-bot.git`
   - `cd 3x-ui-bot`
   - `nano docker-compose.yaml`
     - put your `bot token` after: `BotToken=`
   - put your servers information in `requirements.yaml` file:
     - server ip
     - panel port
     - panel username
     - panel password

     > If you want to add another server, uncomment other lines


   - `sh install.sh`

## ⚠️ Attention
dont use `- (dash)` in user email in panel 