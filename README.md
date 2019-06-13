# なにこれ
* Slackのステータスを変更するんだよぉ

# どうつかうんだ
1. `config.go`にパラメータを入力
    * `OWMApiKey`に**OpenWeatherMap**で発行したAPIKey
    * `SlackUserToken`に**Slack**で発行したトークン（"Bearer xoxp-..."の形で入力）
    * `CityLon CityLat`に天気を取得したい都市の緯度経度（最も近い位置のデータが取得される）
2. ```bash
    go run changeSlackStatus.go config.go owmDataStructure.go
    ```
    で実行
3. ステータスが変わる
    * 絵文字: `その時間の天気の絵文字`
    * テキスト: `その時間の天気 今日の最低気温・最高気温`

# これから
- [x] 天気の対応ができていないのでそれ
- [ ] 天気だけじゃ寂しいので他のモードも作る
- [ ] Bot化して自動で更新できるようにする
