# ベースイメージとして公式のAlpineイメージを使用
FROM golang:1.22.3-alpine AS builder

# ワーキングディレクトリを設定
WORKDIR /app

# go.mod と go.sum をコピーして依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# 実行用の小さなイメージを作成
FROM alpine:latest

# 必要なライブラリをインストール
RUN apk --no-cache add ca-certificates

# 作成したバイナリをコピー
COPY --from=builder /app/main /app/main

# ポートを指定
EXPOSE 8080

# コンテナ起動時に実行されるコマンドを指定
CMD ["/app/main"]
