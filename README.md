
## 構成について
| パッケージ | 説明                                            |
|--------|------------------------------------------------|
| adapter | アプリケーションと外部のを繋ぐインターフェイス |
| driver | フレームワークや外部の技術に依存するコードを実装する |
| entity | ドメインロジックの実装 |
| usecase | アプリケーション固有のビジネスロジックの実装 |
| registry | 各レイヤのDIをする |

## sql boilerの自動生成
```console
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

cd pkg/driver/mysql
sqlboiler -o dto -p dto --no-tests --wipe mysql
```

もしかしたら下の対応が必要かも？？
```console
go get -u github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql
go get -u github.com/volatiletech/sqlboiler/v4
go get -u github.com/volatiletech/null/v8

```
もしかしたら下の対応が必要かも？？
```console
rm $GOPATH/bin/sqlboiler*

```

https://github.com/volatiletech/sqlboiler/issues/774