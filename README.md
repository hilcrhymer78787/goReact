https://github.com/hilcrhymer78787/goReact.git

cd goReact

docker-compose up -d --build

docker-compose exec app bash

cp .env.example .env

go run main.go

もう一つターミナルを開く

cd next

cp .env.example .env

npm install

npm run dev

<!-- node v16.13.2 -->