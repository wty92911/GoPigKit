#!/usr/bin/env sh

# shellcheck disable=SC2164
cd "$(cd "$(dirname "$0")";pwd)"
cd ../

# 复制 tar.gz 文件到远程服务器
scp *.tar.gz myserver:~

# 在远程服务器上执行一系列命令
ssh myserver << 'EOF'
# 进入用户主目录
cd ~
# 停止服务
sh pigkit/bin/server.sh stop

# 备份并解压新文件
rm -rf pigkit_backup && mv pigkit pigkit_backup
tar -xzvf *.tar.gz && rm *.tar.gz

# 进入新解压的目录，并启动 Docker Compose
cd pigkit/ && docker-compose up -d && sh bin/server.sh start
EOF