# Goの公式イメージを使う
FROM golang:1.23 AS builder

# 作業ディレクトリの作成
WORKDIR /app

# Goモジュールのコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# 最終的なイメージ
FROM debian:bookworm-slim 

# ビルド済みアプリケーションをコピー
COPY --from=builder /app/main /main

# アプリケーションの実行
ENTRYPOINT ["/main"]
