git clone https://github.com/hilcrhymer78787/vuego.git

cd vuego

docker-compose up -d --build

docker-compose exec app bash

cp .env.example .env

go run main.go

もう一つターミナルを開く

cd nuxt

npm install

npm run dev

<!-- node v16.13.2 -->