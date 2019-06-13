package main

import (
	"net/http"
	"net/url"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"
	"math"
	"fmt"
	"strconv"
)

type weatherSlackProfile struct {
	StatusText  string `json:"status_text"`
	StatusEmoji string `json:"status_emoji"`
}
type weatherSlackStatus struct {
	Token   string  `json:"token"`
	Profile weatherSlackProfile `json:"profile"`
}

// Map for OWM weather condition
// https://openweathermap.org/weather-conditions
type weatherCondition struct {
	Description string
	Emoji string
}

func main() {
	cityData := url.Values{}
	cityData.Add("lon", CityLon)
	cityData.Add("lat", CityLat)
	cityData.Add("cnt", "8")
	cityData.Add("APPID", OWMApiKey)

	statusData, err := getCityWeather(cityData)
	if err != nil {
		panic(err)
	}

	err = postSlackStatus(*statusData)
	if err != nil {
		panic(err)
	}
}

func getCityWeather(cityData url.Values) (*weatherSlackStatus, error) {
	apiURL := OWMApiURL + "?" + cityData.Encode()
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jsonBytes := ([]byte)(byteArray)
	data := new(OWM5DaysWeatherResponse)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
        return nil, err
    }

	statusData := new(weatherSlackStatus)
	statusData.Token = SlackUserToken
	weatherID := data.List[0].Weather[0].ID
	weatherDetail := WeatherMap[weatherID]
	statusText := weatherDetail.Description
	tempMin := 1000.0
	tempMax := -1000.0
	today := time.Now().Day()
	for _, d := range data.List {
		if time.Unix(d.Dt, 0).Day() == today {
			tempMin = math.Min(tempMin, d.Main.TempMin)
			tempMax = math.Max(tempMax, d.Main.TempMax)
		} else {
			break
		}
	}
	tempMin = tempMin - 273.15
	tempMax = tempMax - 273.15
	statusText += " " + strconv.FormatFloat(tempMin, 'f', 1, 64) + "~" + strconv.FormatFloat(tempMax, 'f', 1, 64) + "℃"

	fmt.Println(statusText)
	statusData.Profile = weatherSlackProfile {
		StatusText: statusText,
		StatusEmoji: weatherDetail.Emoji,
	}
	return statusData, err
}

func postSlackStatus(statusData weatherSlackStatus) error {
	jsonData, err := json.Marshal(statusData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		SlackAPIURL,
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", SlackUserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
