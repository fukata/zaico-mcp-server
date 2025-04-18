# zaico-mcp-server

zaico-mcp-serverは、zaico APIとMark3Labs MCPを統合するためのサーバーアプリケーションです。

## 機能

- 在庫データ
  - 一覧取得

## 必要条件

- Go 1.23.7以上

## インストール

### 方法1: go installを使用する場合

```bash
# 最新バージョンをインストール
go install github.com/fukata/zaico-mcp-server/cmd/zaico-mcp-server@latest

# または特定のバージョンをインストール
go install github.com/fukata/zaico-mcp-server/cmd/zaico-mcp-server@v1.0.0
```

インストール後は、`$GOPATH/bin`に配置されるため、以下のコマンドで実行できます：

```bash
zaico-mcp-server --zaico-api-key <APIキー>
```

### 方法2: バイナリを直接ダウンロードする場合

[リリースページ](https://github.com/fukata/zaico-mcp-server/releases)から、お使いのOSに合わせたバイナリをダウンロードしてください。

#### macOSでの実行方法

macOSで実行する場合、セキュリティ上の理由から以下の手順が必要です：

1. ターミナルを開く
2. ダウンロードしたファイルがあるディレクトリに移動
3. 以下のコマンドを実行：

```bash
# 実行権限を付与
chmod +x zaico-mcp-server-darwin-arm64

# Gatekeeperの制限を解除
xattr -d com.apple.quarantine zaico-mcp-server-darwin-arm64

# 実行
./zaico-mcp-server-darwin-arm64
```

## 使用方法

```bash
zaico-mcp-server --zaico-api-key <APIキー>
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