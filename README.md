# SmartAttendance
NFCタグを用いて出席確認を行うアプリ

# 環境構築
## エラー
エラー文：network job_hunting_network declared as external, but could not be found
エラー内容：コンテナ間をつなぐネットワークがない
解決方法: docker network create job_hunting_network

## airが起動しないとき
docker-compose build
docker-compose up

# SQLファイルの実行方法
docker exec  -it smart_attendance_db bash

mysql -u user -p smart_attendance_database

source docker-entrypoint-initdb.d/01_create_users.sql
source docker-entrypoint-initdb.d/02_create_classes.sql
source docker-entrypoint-initdb.d/03_create_user_class.sql

# デバッグ
## curlによるユーザー登録
### "の前に\をつけてエスケープする
curl -X POST -H "Content-Type: application/json" -d "{\"id\" : \"tanaka\" , \"name\": \"田中\", \"email\": \"tanaka@example.com\", \"password\": \"tanaka\"}" localhost:8080/signup

## curlによるユーザー一覧確認
curl localhost:8080/users

## ユーザーのログイン確認
### {"token":"@@@..."} が帰ってきたら成功
curl -X POST -H "Content-Type: application/json" -d "{\"email\" : \"tanaka@example.com\" , \"password\": \"tanaka\"}" localhost:8080/login

### tokenを渡してAPIを取得する
curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJ0YW5ha2EiLCJuYW1lIjoi55Sw5LitIiwiZXhwIjoxNzI4MTkwNDQ1fQ.8zjpCEQFVmDrI7i7Qi2DS4W9X7jmnciR8ChP5mW0s6c" localhost:8080/api/user