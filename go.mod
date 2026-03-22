// cdktf get が Go バインディング生成時にモジュール名を参照するために必要。
// 実際の Go モジュールは src/hashicorp/google/ と src/hashicorp/googlebeta/ にある。
module github.com/builtbystack/cdktf-provider

go 1.26.0
