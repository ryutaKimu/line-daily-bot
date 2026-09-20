# line-daily-bot

鋼の錬金術師の次回予告ナレーションをランダムに1件取得して、毎朝 LINE に push 送信する Go 製のボットです。GitHub Actions のスケジュール実行で動きます。

## 送信されるメッセージ

```
Episode 第39話
『白昼の夢』

未知の世界への旅立ち
それは新しい物を求める心
人ならば心を止める事はできない
だが、そこには必ず軋轢が生じる
次回、鋼の錬金術師 FULLMETAL ALCHEMIST 第40話『フラスコの中の小人（ホムンクルス）』
正しい者が勝つのではない、壁を乗り越えた者が勝つのだ
```

## しくみ

1. `fetchNarrationAPI()` — [fullmetalapi](https://fullmetalapi.vercel.app/narrations/random) からナレーションを1件取得する
2. `buildMessage()` — エピソード番号・タイトル・ナレーション各行を上記フォーマットに整形する
3. `main()` — LINE Messaging API の push エンドポイントへ送信する

## 必要な環境変数

| 変数 | 内容 |
| --- | --- |
| `LINE_TOKEN` | LINE Messaging API のチャネルアクセストークン |
| `LINE_TO` | 送信先のユーザーID / グループID |

## ローカルでの実行

```sh
LINE_TOKEN=xxxxx LINE_TO=yyyyy go run .
```

## テスト

```sh
go test ./...
```

`TestBuildMessage` は固定値でフォーマットを検証します。`TestFetchNarrationAPI` と `TestFormatNarration` は実際の API を叩くため、ネットワーク接続が必要です。

## 定期実行

`.github/workflows/daily.yml` が毎日 23:00 UTC（8:00 JST）に `go run .` を実行します。`LINE_TOKEN` と `LINE_TO` はリポジトリの Actions secrets に登録してください。`workflow_dispatch` で手動実行もできます。
