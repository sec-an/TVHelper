#!/bin/bash

# INSTALL_PATH='/opt/TVHelper'
VERSION='latest'

if [ ! -n "$2" ]; then
  INSTALL_PATH='/opt/TVHelper'
else
  if [[ $2 == */ ]]; then
    INSTALL_PATH=${2%?}
  else
    INSTALL_PATH=$2
  fi
  if ! [[ $INSTALL_PATH == */TVHelper ]]; then
    INSTALL_PATH="$INSTALL_PATH/TVHelper"
  fi
fi

RED_COLOR='\e[1;31m'
GREEN_COLOR='\e[1;32m'
YELLOW_COLOR='\e[1;33m'
BLUE_COLOR='\e[1;34m'
PINK_COLOR='\e[1;35m'
SHAN='\e[1;33;5m'
RES='\e[0m'
clear

# Get platform
if command -v arch >/dev/null 2>&1; then
  platform=$(arch)
else
  platform=$(uname -m)
fi

ARCH="UNKNOWN"

if [ "$platform" = "x86_64" ]; then
  ARCH=x86_64
elif [ "$platform" = "aarch64" ]; then
  ARCH=arm64
fi

if [ "$(id -u)" != "0" ]; then
  echo -e "\r\n${RED_COLOR}出错了，请使用 root 权限重试！${RES}\r\n" 1>&2
  exit 1
elif [ "$ARCH" == "UNKNOWN" ]; then
  echo -e "\r\n${RED_COLOR}出错了${RES}，一键安装目前仅支持 x86_64和arm64 平台。\r\n其它平台请参考：${GREEN_COLOR}https://github.com/sec-an/TVHelper${RES}\r\n"
  exit 1
elif ! command -v systemctl >/dev/null 2>&1; then
  echo -e "\r\n${RED_COLOR}出错了${RES}，无法确定你当前的 Linux 发行版。\r\n建议手动安装：${GREEN_COLOR}https://github.com/sec-an/TVHelper${RES}\r\n"
  exit 1
else
  if command -v netstat >/dev/null 2>&1; then
    check_port=$(netstat -lnp | grep 16214 | awk '{print $7}' | awk -F/ '{print $1}')
  else
    echo -e "${GREEN_COLOR}端口检查 ...${RES}"
    if command -v yum >/dev/null 2>&1; then
      yum install net-tools -y >/dev/null 2>&1
      check_port=$(netstat -lnp | grep 16214 | awk '{print $7}' | awk -F/ '{print $1}')
    else
      apt-get update >/dev/null 2>&1
      apt-get install net-tools -y >/dev/null 2>&1
      check_port=$(netstat -lnp | grep 16214 | awk '{print $7}' | awk -F/ '{print $1}')
    fi
  fi
fi

CHECK() {
  if [ -f "$INSTALL_PATH/TVHelper" ]; then
    echo "此位置已经安装，请选择其他位置，或使用更新命令"
    exit 0
  fi
  if [ $check_port ]; then
    kill -9 $check_port
  fi
  if [ ! -d "$INSTALL_PATH/" ]; then
    mkdir -p $INSTALL_PATH
  else
    rm -rf $INSTALL_PATH && mkdir -p $INSTALL_PATH
  fi
}

INSTALL() {
  # 下载 TVHelper 程序
  echo -e "\r\n${GREEN_COLOR}下载 TVHelper $VERSION ...${RES}"
  curl -L https://gh-proxy.com/https://github.com/sec-an/TVHelper/releases/latest/download/TVHelper_Linux_$ARCH.tar.gz -o /tmp/TVHelper.tar.gz $CURL_BAR
  tar zxf /tmp/TVHelper.tar.gz -C $INSTALL_PATH/

  if [ -f $INSTALL_PATH/TVHelper ]; then
    echo -e "${GREEN_COLOR} 下载成功 ${RES}"
  else
    echo -e "${RED_COLOR}下载 TVHelper_Linux_$ARCH.tar.gz 失败！${RES}"
    exit 1
  fi

  # 删除下载缓存
  rm -f /tmp/TVHelper*
}

INIT() {
  if [ ! -f "$INSTALL_PATH/TVHelper" ]; then
    echo -e "\r\n${RED_COLOR}出错了${RES}，当前系统未安装 TVHelper\r\n"
    exit 1
  fi

  # 创建 systemd
  cat >/etc/systemd/system/TVHelper.service <<EOF
[Unit]
Description=TVHelper service
Wants=network.target
After=network.target network.service

[Service]
Type=simple
WorkingDirectory=$INSTALL_PATH
ExecStart=$INSTALL_PATH/TVHelper
KillMode=process
Restart=always

[Install]
WantedBy=multi-user.target
EOF

  # 添加开机启动
  systemctl daemon-reload
  systemctl enable TVHelper >/dev/null 2>&1
  systemctl restart TVHelper
}

SUCCESS() {
  clear
  echo "TVHelper 安装成功！"
  echo -e "\r\n访问地址：${GREEN_COLOR}http://YOUR_IP:16214/${RES}\r\n"

  echo -e "请详细阅读：${GREEN_COLOR}$INSTALL_PATH/configs/config/sample.json${RES}"
  
  echo
  echo -e "查看状态：${GREEN_COLOR}systemctl status TVHelper${RES}"
  echo -e "启动服务：${GREEN_COLOR}systemctl start TVHelper${RES}"
  echo -e "重启服务：${GREEN_COLOR}systemctl restart TVHelper${RES}"
  echo -e "停止服务：${GREEN_COLOR}systemctl stop TVHelper${RES}"
  echo -e "\r\n温馨提示：如果端口无法正常访问，请检查 \033[36m服务器安全组、本机防火墙、TVHelper状态\033[0m"
  echo
}

UNINSTALL() {
  echo -e "\r\n${GREEN_COLOR}卸载 TVHelper ...${RES}\r\n"
  echo -e "${GREEN_COLOR}停止进程${RES}"
  systemctl disable TVHelper >/dev/null 2>&1
  systemctl stop TVHelper >/dev/null 2>&1
  echo -e "${GREEN_COLOR}清除残留文件${RES}"
  rm -rf $INSTALL_PATH /etc/systemd/system/TVHelper.service
  # 兼容之前的版本
  rm -f /lib/systemd/system/TVHelper.service
  systemctl daemon-reload
  echo -e "\r\n${GREEN_COLOR}TVHelper 已在系统中移除！${RES}\r\n"
}

UPDATE() {
  if [ ! -f "$INSTALL_PATH/TVHelper" ]; then
    echo -e "\r\n${RED_COLOR}出错了${RES}，当前系统未安装 TVHelper\r\n"
    exit 1
  else
    echo
    echo -e "${GREEN_COLOR}停止 TVHelper 进程${RES}\r\n"
    systemctl stop TVHelper
    # 备份 TVHelper 二进制文件，供下载更新失败回退
    cp $INSTALL_PATH/TVHelper /tmp/TVHelper.bak
    echo -e "${GREEN_COLOR}下载 TVHelper $VERSION ...${RES}"
    curl -L https://gh-proxy.com/https://github.com/sec-an/TVHelper/releases/latest/download/TVHelper_Linux_$ARCH.tar.gz -o /tmp/TVHelper.tar.gz $CURL_BAR
    tar zxf /tmp/TVHelper.tar.gz -C $INSTALL_PATH/
    if [ -f $INSTALL_PATH/TVHelper ]; then
      echo -e "${GREEN_COLOR} 下载成功 ${RES}"
    else
      echo -e "${RED_COLOR}下载 TVHelper_Linux_$ARCH.tar.gz 出错，更新失败！${RES}"
      echo "回退所有更改 ..."
      mv /tmp/TVHelper.bak $INSTALL_PATH/TVHelper
      systemctl start TVHelper
      exit 1
    fi
    echo -e "\r\n${GREEN_COLOR}启动 TVHelper 进程${RES}"
    systemctl start TVHelper
    echo -e "\r\n${GREEN_COLOR}TVHelper 已更新到最新稳定版！${RES}\r\n"
    # 删除临时文件
    rm -f /tmp/TVHelper*
  fi
}

# CURL 进度显示
if curl --help | grep progress-bar >/dev/null 2>&1; then # $CURL_BAR
  CURL_BAR="--progress-bar"
fi

# The temp directory must exist
if [ ! -d "/tmp" ]; then
  mkdir -p /tmp
fi

# Fuck bt.cn (BT will use chattr to lock the php isolation config)
chattr -i -R $INSTALL_PATH >/dev/null 2>&1

if [ "$1" = "uninstall" ]; then
  UNINSTALL
elif [ "$1" = "update" ]; then
  UPDATE
elif [ "$1" = "install" ]; then
  CHECK
  INSTALL
  INIT
  if [ -f "$INSTALL_PATH/TVHelper" ]; then
    SUCCESS
  else
    echo -e "${RED_COLOR} 安装失败${RES}"
  fi
else
  echo -e "${RED_COLOR} 错误的命令${RES}"
fi