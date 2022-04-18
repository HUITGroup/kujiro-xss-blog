# API仕様

## POST

### POST /post **ヘッダーにContent-Type: application/json が必須**

記事を投稿する。

```json
{
    "title": "タイトル",
    "conent": "hello world!"
}
 ```

### POST /post/:post_id/comment **ヘッダーにContent-Type: application/json が必須**

特定の記事(post_id)に対してコメントを投稿する。

```json
{
    "name": "hoge",
    "content": "hello comment!"
}
```

## GET

### GET /post/list?limit=2

記事リストの取得する。
クエリのlimitは件数制限(記事数<limitの場合、記事数分のみ返る)

```json
[{
    "post_id": 1234,
    "title": "タイトル",
    "content": "hello world! hi...",
    "timestamp": "2022-04-06T12:57:42+09:00"
},{
    "post_id": 1235,
    "title": "タイトル",
    "content": "hello world! hi...",
    "timestamp": "2022-04-06T12:57:42+09:00"
}]
```

### GET /post/:post_id

特定の記事(post_id)を取得する。

```json
{
    "post_id": 1234,
    "title": "タイトル",
    "content": "hello world! hi...",
    "timestamp": "2022-04-06T12:57:42+09:00"
}
```

### GET /post/:post_id/comments?limit=2

特定の記事(post_id)のコメントを取得する。
クエリのlimitは件数制限

```json
[{
    "comment_id": 5435235,
    "name": "hoge",
    "content": "test comment1",
    "post_id": 1234,
    "timestamp": "2022-04-06T12:57:42+09:00"
},{
    "comment_id": 5435236,
    "name": "fuga",
    "content": "test comment2",
    "post_id": 1234,
    "timestamp": "2022-04-06T12:57:42+09:00"
}]
```
