#!/bin/bash/

# variables
red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

# helpers
insert_in_yaml() {
    id=$(head /dev/urandom | tr -dc 'a-zA-Z0-9' | head -c 8)
    local ip="$1"
    local port="$2"
    local username="$3"
    local password="$4"

    result=$(yq --yaml-output ".servers += [{"id": \""$id"\", "ip": \""$ip"\", "port": \""port"\", "username": \""$username"\", "password": \""$password"\"}]" requirements.yaml)
    echo "$result" > requirements.yaml
}
update_yaml() {
    local id="$1"
    local new_ip="$2"
    local new_port="$3"
    local new_username="$4"
    local new_password="$5"

    command=$(yq --yaml-output \
      ".servers |= map(if .id == \"$id\" then
                        . + 
                        {
                          \"ip\": $(if [ -n "$new_ip" ]; then echo "\"$new_ip\""; else echo ".ip"; fi),
                          \"port\": $(if [ -n "$new_port" ]; then echo "$new_port"; else echo ".port"; fi),
                          \"username\": $(if [ -n "$new_username" ]; then echo "\"$new_username\""; else echo ".username"; fi),
                          \"password\": $(if [ -n "$new_password" ]; then echo "\"$new_password\""; else echo ".password"; fi)
                        }
                      else
                        .
                      end)" \
      requirements.yaml)
     echo "$command" > requirements.yaml
}

# menu functions
set_bot_token() {
    echo && read -p "Please enter your telegram bot token: " token
    command=$(yq --yaml-output '.services."3x-ui-bot".environment[0] = "BotToken='$token'"' docker-compose.yaml)
    echo "$command" > docker-compose.yaml
    echo ${green}"[ Bot token set ]"
}
get_servers() {
	cat requirements.yaml
}
define_new_server() {
    echo "Please enter your new server information: "
    echo -n && read -p "server ip: " ip
    echo -n && read -p "server port: " port
    echo -n && read -p "server usernamse: " username
    echo -n && read -p "server password: " password

    insert_in_yaml $ip $port $username $password
    echo ${green}"[ New server added ]"
}
update_server_info() {
    echo && read -p "server id: " id
    echo "Please enter your new server information: "
    echo -n && read -p "server ip: " ip
    echo -n && read -p "server port: " port
    echo -n && read -p "server usernamse: " username
    echo -n && read -p "server password: " password

    update_yaml $id $ip $port $username $password
    echo ${green}"[ Server info updated ]"
}
remove_server() {
    echo && read -p "Please enter your server id: " id

    command='yq --yaml-output "del(.servers[] | select(.id == \""$id"\"))" requirements.yaml'
    result="$(eval "$command")"
    echo "$result" > requirements.yaml
}
remove_all_servers() {
    content="servers:"
    echo $content > requirements.yaml
    echo ${green}"[ All servers removed ]"
}

show_menu() {
    echo "
  ${green}3X-ui Panel Management Script${plain}
  ${green}0.${plain} Exit Script
————————————————
  ${green}1.${plain} Set Telegram bot token
  ${green}2.${plain} Get all your servers
  ${green}3.${plain} Define new server
  ${green}4.${plain} Update server info
  ${green}5.${plain} Remove server
  ${green}6.${plain} Remove all servers
"
    echo -n && read -p "Please enter your selection [0-22]: " num

    case "${num}" in
    0)
        exit 0
        ;;
    1)
        set_bot_token
        ;;
    2)
	get_servers
	;;
    3)
        define_new_server
	;;
    4)
	update_server_info
	;;
    5)
	remove_server
	;;
    6)
	remove_all_servers
	;;
    *)
        LOGE "Please enter the correct number [0-22]"
        ;;
    esac
}

while true; do
    show_menu
done
