
###
 # @Author: Vincent Yang
 # @Date: 2024-04-26 22:06:48
 # @LastEditors: Vincent Yang
 # @LastEditTime: 2024-04-26 22:09:23
 # @FilePath: /UpdateCompose/install.sh
 # @Telegram: https://t.me/missuo
 # @GitHub: https://github.com/missuo
 # 
 # Copyright © 2024 by Vincent, All Rights Reserved. 
### 
install_update-compose() {
  if [[ $EUID -ne 0 ]]; then
    echo -e "${red}Error: This script must be run as root.${plain}" 
    exit 1
  fi

  last_version=$(curl -Ls "https://api.github.com/repos/missuo/UpdateCompose/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
  if [[ ! -n "$last_version" ]]; then
    echo -e "${red}Failed to detect UpdateCompose version, probably due to exceeding Github API limitations.${plain}"
    exit 1
  fi

  echo -e "UpdateCompose latest version: ${last_version}, Start install..."
  wget -q -N --no-check-certificate -O /usr/bin/update-compose https://github.com/missuo/UpdateCompose/releases/download/${last_version}/update-compose-linux-amd64
  chmod +x /usr/bin/update-compose
  echo -e "UpdateCompose installed successfully! Just run update-compose to use it."
}

install_update-compose