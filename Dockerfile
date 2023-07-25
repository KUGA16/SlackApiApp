# ベースイメージ
FROM golang:latest
# コンテナ内にディレクトリを作成
RUN mkdir /go/src/slackApiApp/
# ワーキングディレクトリの設定
WORKDIR /go/src/slackApiApp/
# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY ./ /go/src/slackApiApp/
