#!/usr/bin/env bash

API_URL="https://localhost"

# 記事投稿
curl -X POST -H 'Content-Type: application/json' -d '{"title":"タイトル1", "content":"コンテンツ1"}' "${API_URL}/post"
curl -X POST -H 'Content-Type: application/json' -d '{"title":"タイトル2", "content":"コンテンツ2"}' "${API_URL}/post"
curl -X POST -H 'Content-Type: application/json' -d '{"title":"タイトル3", "content":"コンテンツ3"}' "${API_URL}/post"

# 記事リスト取得
curl -X GET "${API_URL}/post/list?limit=1"
curl -X GET "${API_URL}/post/list?limit=2"
curl -X GET "${API_URL}/post/list?limit=3"
curl -X GET "${API_URL}/post/list?limit=4"

# 記事取得
curl -X GET "${API_URL}/post/1"
curl -X GET "${API_URL}/post/2"
curl -X GET "${API_URL}/post/10"

# コメント投稿
curl -X POST -H 'Content-Type: application/json' -d '{"name":"名前1", "content":"コメント1"}' "${API_URL}/post/1/comment"
curl -X POST -H 'Content-Type: application/json' -d '{"name":"名前2", "content":"コメント2"}' "${API_URL}/post/1/comment"
curl -X POST -H 'Content-Type: application/json' -d '{"name":"名前3", "content":"コメント3"}' "${API_URL}/post/1/comment"

# コメント取得
curl -X GET "${API_URL}/post/1/comments?limit=1"
curl -X GET "${API_URL}/post/1/comments?limit=2"
curl -X GET "${API_URL}/post/1/comments?limit=3"
curl -X GET "${API_URL}/post/1/comments?limit=5"
