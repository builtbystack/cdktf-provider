#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

MODULE=github.com/builtbystack/cdktf-provider
GO_VERSION=1.26.0

# Go モジュールの 500MB サイズ上限を超えるため google と google_beta を別モジュールに分離している。
# cdktf get は go.mod のモジュール名を参照するためルートの go.mod が必要。
# 生成後、各プロバイダディレクトリに go.mod がなければ作成し go mod tidy で依存を解決する。

# cdktf get が go.mod を削除する場合があるため、存在しなければ再生成してから go mod tidy する
setup_go_mod() {
  if [[ ! -f "$1/go.mod" ]]; then
    go mod init -C "$1" "${MODULE}/$1"
    go mod edit -C "$1" -go="$GO_VERSION"
  fi
  go mod tidy -C "$1"
}

# プロバイダのスキーマが大きいためメモリ上限を引き上げないと out of memory になる
NODE_OPTIONS="--max-old-space-size=4096" npx cdktf-cli get

setup_go_mod src/hashicorp/google
setup_go_mod src/hashicorp/google_beta
