# ベースとなるDockerイメージ指定
FROM golang:latest

RUN curl -sL https://deb.nodesource.com/setup_15.x | bash -
RUN apt-get install -y nodejs
RUN npm install npm@latest -g

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/go

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/go