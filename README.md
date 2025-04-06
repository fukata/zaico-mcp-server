# Zaico MCP Server

Zaico MCP Serverは、Zaico APIとMark3Labs MCPを統合するためのサーバーアプリケーションです。

## 機能

- Zaico APIとの連携
- Mark3Labs MCPとの統合
- 設定ファイルによる柔軟な設定管理
- ログ出力機能

## 必要条件

- Go 1.23.7以上
- 必要な環境変数の設定

## インストール

```bash
go get github.com/fukata/zaico-mcp-server
```

## 使用方法

1. 設定ファイルを作成します（`config.yaml`）
2. サーバーを起動します：

```bash
./zaico-mcp-server
```

## 設定

設定は以下の方法で行えます：

- 環境変数
- 設定ファイル（YAML）
- コマンドライン引数

## ディレクトリ構造

```
.
├── cmd/          # メインアプリケーション
├── pkg/          # 公開パッケージ
```

## ライセンス

このプロジェクトは[MITライセンス](LICENSE.txt)の下で公開されています。 