package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type EasyWeather struct {
	CityInfo    CityInfo        `json:"cityInfo"`
	WeatherData EasyWeatherData `json:"data"`
}

type EasyWeatherData struct {
	Forecasts []EForecast `json:"forecast"`
	Yesterday EForecast   `json:"yesterday"`
}

type EForecast struct {
	Week string `json:"week"`
	High int    `json:"high"`
	Low  int    `json:"low"`
	Type string `json:"type"`
}

func ForecastToEasyForecast(forecast Forecast) EForecast {
	var eForecast EForecast
	eForecast.Week = forecast.Week
	eForecast.Type = forecast.Type
	// log.Debug("high0:" + forecast.High)
	high := strings.Replace(forecast.High, "高温 ", "", -1)
	// log.Debug("high1:" + high)
	high = strings.Replace(high, "℃", "", -1)
	// log.Debug("high2:" + high)
	eForecast.High, _ = strconv.Atoi(high)
	// log.Debugf("F high:%d\n", eForecast.High)
	low := strings.Replace(forecast.Low, "低温 ", "", -1)
	low = strings.Replace(low, "℃", "", -1)
	eForecast.Low, _ = strconv.Atoi(low)
	return eForecast
}

func TranWeatherToEWeather(weather *Weather) *EasyWeather {
	easyWeather := &EasyWeather{}
	easyWeather.CityInfo = weather.CityInfo
	for _, value := range weather.WeatherData.Forecasts {
		eForecast := ForecastToEasyForecast(value)
		easyWeather.WeatherData.Forecasts = append(easyWeather.WeatherData.Forecasts, eForecast)
	}
	easyWeather.WeatherData.Yesterday = ForecastToEasyForecast(weather.WeatherData.Yesterday)
	return easyWeather
}

type Weather struct {
	Status      int         `json:"status"`
	Date        string      `json:"date"`
	Time        string      `json:"time"`
	CityInfo    CityInfo    `json:"cityInfo"`
	WeatherData WeatherData `json:"data"`
}

type WeatherData struct {
	Shidu     string     `json:"shidu"`
	Pm25      float64    `json:"pm25"`
	Pm10      float64    `json:"pm10"`
	Quality   string     `json:"quality"`
	Wendu     string     `json:"wendu"`
	Ganmao    string     `json:"ganmao"`
	Forecasts []Forecast `json:"forecast"`
	Yesterday Forecast   `json:"yesterday"`
}

type CityInfo struct {
	City       string `json:"city"`
	CityKey    string `json:"citykey"`
	Parent     string `json:"parent"`
	UpdateTime string `json:"updateTime"`
}
type Forecast struct {
	Date    string `json:"date"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Ymd     string `json:"ymd"`
	Week    string `json:"week"`
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
	Api     int    `json:"api"`
	Fx      string `json:"fx"`
	Fl      string `json:"fl"`
	Type    string `json:"type"`
	Notice  string `json:"notice"`
}

//GetWeather 获取城市15天天气
func GetWeather(cityCode string) (*Weather, error) {

	url := "http://t.weather.sojson.com/api/weather/city/" + cityCode
	weather := &Weather{}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(out, weather); err != nil {
		return nil, err
	}

	if weather.Status != 200 {
		return nil, errors.New("get weather error!")
	}

	return weather, nil
}
