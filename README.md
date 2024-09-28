# atcoder-go

[AtCoder](https://atcoder.jp/?lang=ja)のgo言語による解答を記録するリポジトリ。

テンプレートからディレクトリとファイルを生成するには、以下のコマンドを実行する：

```
$ bin/generate -c abc100 -p a,b,c,d
$ tree contests/abc100           
contests/abc100
├── a
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
├── b
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
├── c
│   ├── go.mod
│   ├── main.go
│   └── main_test.go
└── d
    ├── go.mod
    ├── main.go
    └── main_test.go

5 directories, 12 files
```
