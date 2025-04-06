# zaico-mcp-server

zaico-mcp-serverは、zaico APIとMark3Labs MCPを統合するためのサーバーアプリケーションです。

## 機能

- 在庫データ
  - 一覧取得
  - 個別取得
  - 作成
  - 更新
  - 削除

## 必要条件

- Go 1.23.7以上

## インストール

```bash
go get github.com/fukata/zaico-mcp-server
```

## 使用方法

```bash
./zaico-mcp-server --zaico-api-key <APIキー>
```

## 設定

設定は以下の方法で行えます：

- コマンドライン引数

## ディレクトリ構造

```
.
├── cmd/          # メインアプリケーション
├── pkg/          # 公開パッケージ
```

## ライセンス

このプロジェクトは[MITライセンス](LICENSE.txt)の下で公開されています。 