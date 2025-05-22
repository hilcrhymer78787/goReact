https://github.com/hilcrhymer78787/goReact.git

cd goReact

docker-compose up -d --build

docker-compose exec app bash

cp .env.example .env

air <!-- go run main.go と同じ意味-->

もう一つターミナルを開く

cd next

cp .env.example .env

nvm use v16.13.2

yarn install

yarn dev

<!-- node v16.13.2 -->
