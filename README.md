# upsider-coding-test

## 開発

```bash
make start-dev
```

起動したコンテナにアタッチし、コンテナ内で開発・テストを行ってください。  

```bash
# 必要な開発ツールのインストール
make setup

# DBマイグレーション
make migration

# DBマイグレーションのロールバック
make rollback
```

## テスト

### e2e

下記コマンドでAPI起動

```bash
make serve
```

VSCodeの拡張機能であるREST Clientをインストールする。  
[.http/api.http](.http/api.http)を開き、リクエストを送る。
