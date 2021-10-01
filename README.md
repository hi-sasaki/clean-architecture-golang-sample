
## 構成について
| パッケージ | 説明                                            |
|--------|------------------------------------------------|
| adapter | アプリケーションと外部のを繋ぐインターフェイス |
| driver | フレームワークや外部の技術に依存するコードを実装する |
| entity | ドメインロジックの実装 |
| usecase | アプリケーション固有のビジネスロジックの実装 |
| registry | 各レイヤのDIをする |

## 環境変数の設定
```sh
cp envrc.sample .envrc
direnv edit .
direnv allow .
```
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

curl -X POST -d '{"password":"abcd","first_name":"sasaki","last_name":"hiroto"}'  http://localhost:8080/user
