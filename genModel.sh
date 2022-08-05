#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./genModel

# mysql数据库配置
#host=127.0.0.1
#port=3306
#dbname=$1
#username=root
#passwd=123456


echo "开始创建库：$dbname 的表：$2"
#goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero


# postgresql数据库配置
host=127.0.0.1
port=5432
dbname=$1
username=postgres
passwd=123456

echo "开始创建库：$dbname 的表：$2"
# postgres The data source of database,like "postgres://root:password@127.0.0.1:5432/database?sslmode=disable"
# goctl model pg datasource --url="postgres://postgres:123456@127.0.0.1:5432/go-zero-demo?sslmode=disable" --table=user_s1 --schema=public --dir
goctl model pg datasource -url="postgres://${username}:${passwd}@${host}:${port}/${dbname}?sslmode=disable" -table="${tables}" -schema=public -dir="${modeldir}" -cache=true --style=goZero
