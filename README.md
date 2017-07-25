# slack_client
go製のSlackWebHook用クライアントです。
環境変数でwebhookのURLを設定して、cronなどで定期的なメッセージを特定のチャンネルに送信するなどの用途で利用することを想定しています。

# 使い方

## 環境変数設定

このクライアントを実行するには環境変数として `SLACKURL` を設定してから実行するようにしてください。

例：`export SLACKURL=https://hooks.slack.com/services/hogehoge`

## 実行例

goでビルド後の実行は以下のようにします。

`slack_client 通知したいメッセージ`
