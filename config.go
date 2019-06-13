package main

// OpenWeatherMap API info

// CityLon 都市経度
const CityLon = "139.700338" // Tokyo Shinjuku
// CityLat 都市緯度
const CityLat = "35.686427"
// OWMApiURL 3時間ごと、5日分の天気
const OWMApiURL = "http://api.openweathermap.org/data/2.5/forecast"
// その時間の天気
// const OWMApiURL = "http://api.openweathermap.org/data/2.5/weather"

// OWMApiKey API Token. 取得する必要あり
const OWMApiKey = ""

// Slack API info

// SlackAPIURL API URL
const SlackAPIURL = "https://slack.com/api/users.profile.set"
// SlackUserToken API Token. 取得する必要あり
const SlackUserToken = "Bearer xoxp-"
