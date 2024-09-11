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

# (optional) DBマイグレーションのロールバック
make rollback
```

### ディレクトリ

| ディレクトリ | 概要 |
|:---|:---|
| .http/ | VSCode REST Clientで使用するリクエストファイルの格納 |
| .script/ | 開発の補助スクリプトの格納 |
| cmd/ | エントリーポイントのmain.goの格納 |
| db/ | マイグレーションファイルなどの格納 |
| docs/ | 設計の図などのドキュメントの格納 |
| domain/ | ドメインモデルの格納 |
| infrastructure/ | 外部知識や技術に関する処理の格納 |
| mock/ | `make gen-mock`で生成されるmockの格納、開発者が直接編集することは基本的になし |
| shared/ | domain~infrastructure共通で使用されうる処理の格納 |
| test/ | テストファイルの格納 |
| usecase/ | ユースケースの格納 |

### 注意点

- 小数を含みうる値（金額・手数料率・税率）について、ドメイン内ではdecimalで扱い、ドメイン外では正確性を保つために文字列で扱っています。float64は使用しないでください。
- DIにWireを使用しており、interfaceの実装があるディレクトリにprovider.goという名前のファイルでinterfaceと実装のBindをしつつProvideしています。  
provider.goに追加をした際は`$ make gen-wire`でwire_gen.goを再生成してください。
- interfaceのmockは`$ make gen-mock`で生成してください。

## テスト

現状はdatabaseをテスト時に新規作成・テスト完了したら削除、というようになっておらず1つのdatabaseを使用してしまっていることに留意してください。

### UT, e2e

下記コマンドでUT,e2eテスト含めて実行

```bash
make test
```

### 手動でのAPI起動と任意リクエスト

下記コマンドでAPI起動

```bash
make serve
```

VSCodeの拡張機能であるREST Clientをインストールする。  
[.http/api.http](.http/api.http)を開き、リクエストを送る。
