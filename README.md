# kujiro-xss-blog

2022/04/08 に行われた HUIT 新歓の 「誰ですか こんなところにXSS仕掛けたのは」(kujiro) で使用された XSS デモ用アプリです。

## 事前準備

1. docker と docker-compose のインストール

  それぞれの環境に合わせてインストールをお願いします。

  docker compose (compose v2) でも大丈夫です。

  その場合は `docker-compose` を `docker compose` に読み替えてください。

1. `git clone https://github.com/HUITGroup/kujiro-xss-blog.git && cd kujiro-xss-blog`

## ローカルで遊ぶ場合

1. docker-compose で構築

  ```bash
  docker-compose -f local.compose.yml up
  ```

1. ブラウザで `localhost:8080` に接続して遊ぶ

## サーバに構築する場合

1. SSLの証明書と秘密鍵を [/nginx/ssl](/nginx/ssl) ディレクトリに置く
1. [sample.env](/sample.env) を .env にリネームして内容を記入する (リバースプロキシとなるnginxの設定)
1. docker-compose で構築

  ```bash
  docker-compose -f server.compose.yml up
  ```

1. https://<サーバアドレス> で接続して遊ぶ

  もし接続できない場合は、ファイアウォールやログなどを確認してください。
