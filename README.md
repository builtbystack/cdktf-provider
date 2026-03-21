# cdktf-provider

[CDKTF](https://developer.hashicorp.com/terraform/cdktf) (CDK for Terraform) の Terraform provider バインディングを管理するリポジトリ。
`cdktf get` で生成したプロバイダを Go パッケージとして提供する。

## 利用方法

`src/hashicorp/` 配下にプロバイダごと・リソースごとのパッケージが生成されている。

```sh
go get github.com/builtbystack/cdktf-provider@latest
```

```go
import (
	"github.com/builtbystack/cdktf-provider/src/hashicorp/google/cloudrunv2service"
)
```

## コード生成

前提条件: Node.js, Go

```sh
make
```

## プロバイダのアップグレード

1. `cdktf.json` の `terraformProviders` でバージョン制約を更新する
2. `make` を実行
