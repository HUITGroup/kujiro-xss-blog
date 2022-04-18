# kujiro-xss-frontend

2022/04/08 に行われた HUIT 新歓の 「誰ですか こんなところにXSS仕掛けたのは」(kujiro) で使用された XSS デモ用アプリです。

docker compose を使用して構築します。

## 事前準備

1. SSLの証明書(ssl.crt)と秘密鍵(ssl.key)を [/nginx/ssl](/nginx/ssl) ディレクトリに置く
1. [sample.env](/sample.env) を .env にリネームして内容を記入する

## 構築

```bash
docker compose up -d
```
