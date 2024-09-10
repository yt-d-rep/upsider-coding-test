# upsider-coding-test

## 開発

### 環境構築

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

### 注意点

- DIにWireを使用しており、interfaceの実装があるディレクトリにprovider.goという名前のファイルでinterfaceと実装のBindをしつつProvideしています。  
provider.goに追加をした際は`$ make gen-wire`でwire_gen.goを再生成してください。
- interfaceのmockは`$ make gen-mock`で生成してください。

## テスト

### UT

下記コマンドで実行

```bash
make test
```

### e2e

下記コマンドでAPI起動

```bash
make serve
```

VSCodeの拡張機能であるREST Clientをインストールする。  
[.http/api.http](.http/api.http)を開き、リクエストを送る。
