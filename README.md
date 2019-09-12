# addPreifx

## 概要
ファイル名に対して任意の接頭辞を付与します。
また、ファイル名の先頭から任意の数だけ文字を削除します。

## 使い方
```bash
go build
```
### 接頭辞を付与する
- d: 対象とするディレクトリ
- t: 対象とするファイルの拡張子
- p: 接頭辞として付与する文字列

```bash
./addPrefix -d=testdata/testDir01/ -p=test- -t=.txt
```

### 先頭から文字列を削除する
引数に `del` を指定する

オプション
- d: 対象とするディレクトリ
- t: 対象とするファイルの拡張子
- n: 削除する文字数（先頭から数えた個数）
```bash
./addPrefix -d=testdata/testDir01/ -n=1 -t=.txt del
```
