FROM ubuntu:latest

# set timezone
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# compile and install vim
RUN apt-get update -y \
    && apt-get install vim curl git -y

FROM golang:latest
# コンテナ作業ディレクトリの変更

# # ホストOSの ./go の中身を作業ディレクトリに追加
# ADD ./src .

# # モジュール管理のファイルをコピー
# COPY ./src .
# COPY ./src .
# RUN GO111MODULE=off go get github.com/cosmtrek/air
# RUN go get -u github.com/cosmtrek/air
# # 外部パッケージのダウンロード
ENV GO111MODULE=on

# ENV  GO111MODULE=on
RUN mkdir -p /go/src/magafagafa/go-app
ENV ROOT=/go/src/magafagafa/go-app
WORKDIR ${ROOT}


RUN go mod init github.com/magafagafa/go-app
# RUN go install github.com/pilu/fresh
# RUN go get github.com/labstack/echo/v4
# RUN go get github.com/labstack/echo/v4/middleware
RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air
# ADD . /go/src/magafagafa/go-app
# CMD [ "realize", "start", "--run" ]
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]
SHELL ["/bin/bash", "-c"]

CMD ["air", "-c", ".air.toml"]



