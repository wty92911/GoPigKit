#!/usr/bin/env sh

# 启动 ./server.sh start
# 停止 ./server.sh stop
# 重启 ./server.sh restart

service_name=pigkit
action=$1
pid=$(ps axu | grep "./${service_name} -c" | grep -v grep | awk '{print $2}')

start() {
  # shellcheck disable=SC2164
  bin=$(cd "$(dirname "$0")"; pwd)
  cd ${bin} || exit
  mkdir -p ../log/
  GODEBUG=gctrace=1 nohup ./${service_name}  >> ../log/stdout_${service_name} 2>>../log/stderr_${service_name} &
  echo "$service_name start success."
}

if [ "${action}" = "start" ]; then
  if [ "$pid" != "" ]; then
    echo "[ERROR] Service $service_name is already running"
    exit 1
  fi
  start

elif [ "${action}" = "stop" ]; then
  if [ "$pid" = "" ]; then
    echo "[ERROR] Service ${service_name} is not running"
    exit 1
  fi
  kill -9 ${pid}
  echo "$service_name stopped. kill $pid: $?"

elif [ "${action}" = "restart" ]; then
  if [ "$pid" = "" ]; then
    start
    exit 0
  fi
  kill -9 ${pid}
  echo "$service_name kill $pid: $?"
  start
fi