package wirepod_ttr

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/kercre123/wire-pod/chipper/pkg/logger"
	"github.com/kercre123/wire-pod/chipper/pkg/vars"
)

func getWeather(location string, botUnits string, hoursFromNow int) (string, string, string, string, string, string) {
	var weatherEnabled bool
	var condition string
	var is_forecast string
	var local_datetime string
	var speakable_location_string string
	var temperature string
	var temperature_unit string
	weatherAPIEnabled := vars.APIConfig.Weather.Enable
	weatherAPIKey := vars.APIConfig.Weather.Key
	weatherAPIUnit := vars.APIConfig.Weather.Unit
	weatherAPIProvider := vars.APIConfig.Weather.Provider
	if weatherAPIEnabled && weatherAPIKey != "" {
		weatherEnabled = true
		logger.Println("Weather API enabled")
	} else {
		weatherEnabled = false
		logger.Println("Weather API not enabled, using placeholder")
		if weatherAPIEnabled && weatherAPIKey == "" {
			logger.Println("Weather API enabled, but Weather API key not set")
		}
	}
	if weatherEnabled {
		if botUnits != "" {
			if botUnits == "F" {
				logger.Println("Weather units set to F")
				weatherAPIUnit = "F"
			} else if botUnits == "C" {
				logger.Println("Weather units set to C")
				weatherAPIUnit = "C"
			}
		} else if weatherAPIUnit != "F" && weatherAPIUnit != "C" {
			logger.Println("Weather API unit not set, using F")
			weatherAPIUnit = "F"
		}
	}

	if weatherEnabled {
		if weatherAPIProvider == "weatherapi.com" {
			params := url.Values{}
			params.Add("key", weatherAPIKey)
			params.Add("q", location)
			params.Add("aqi", "no")
			url := "http://api.weatherapi.com/v1/current.json"
			resp, err := http.PostForm(url, params)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			weatherResponse := string(body)
			var weatherAPICladMap weatherAPICladStruct
			mapPath := ""
			if runtime.GOOS == "android" || runtime.GOOS == "ios" {
				mapPath = vars.AndroidPath + "/static/weather-map.json"
			} else {
				mapPath = "./weather-map.json"
			}
			jsonFile, _ := os.ReadFile(mapPath)
			json.Unmarshal(jsonFile, &weatherAPICladMap)
			var weatherStruct weatherAPIResponseStruct
			json.Unmarshal([]byte(weatherResponse), &weatherStruct)
			var matchedValue bool
			for _, b := range weatherAPICladMap {
				if b.APIValue == weatherStruct.Current.Condition.Text {
					condition = b.CladType
					logger.Println("API Value: " + b.APIValue + ", Clad Type: " + b.CladType)
					matchedValue = true
					break
				}
			}
			if !matchedValue {
				condition = weatherStruct.Current.Condition.Text
			}
			is_forecast = "false"
			local_datetime = weatherStruct.Current.LastUpdated
			speakable_location_string = weatherStruct.Location.Name
			if weatherAPIUnit == "C" {
				temperature = strconv.Itoa(int(weatherStruct.Current.TempC))
				temperature_unit = "C"
			} else {
				temperature = strconv.Itoa(int(weatherStruct.Current.TempF))
				temperature_unit = "F"
			}
		} else if weatherAPIProvider == "openweathermap.org" {
			// First use geocoding api to convert location into coordinates
			// E.G. http://api.openweathermap.org/geo/1.0/direct?q={city name},{state code},{country code}&limit={limit}&appid={API key}
			url := "http://api.openweathermap.org/geo/1.0/direct?q=" + url.QueryEscape(location) + "&limit=1&appid=" + weatherAPIKey
			resp, err := http.Get(url)
			if err != nil {
				logger.Println(err)
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			geoCodingResponse := string(body)

			var geoCodingInfoStruct []openWeatherMapAPIGeoCodingStruct

			err = json.Unmarshal([]byte(geoCodingResponse), &geoCodingInfoStruct)
			if err != nil {
				logger.Println(err)
				logger.Println("Geolocation API error: " + geoCodingResponse)
			}
			if len(geoCodingInfoStruct) == 0 {
				logger.Println("Geo provided no response.")
				condition = "undefined"
				is_forecast = "false"
				local_datetime = "test"              // preferably local time in UTC ISO 8601 format ("2022-06-15 12:21:22.123")
				speakable_location_string = location // preferably the processed location
				temperature = "120"
				temperature_unit = "C"
				return condition, is_forecast, local_datetime, speakable_location_string, temperature, temperature_unit
			}
			Lat := fmt.Sprintf("%f", geoCodingInfoStruct[0].Lat)
			Lon := fmt.Sprintf("%f", geoCodingInfoStruct[0].Lon)

			logger.Println("Lat: " + Lat + ", Lon: " + Lon)
			logger.Println("Name: " + geoCodingInfoStruct[0].Name)
			logger.Println("Country: " + geoCodingInfoStruct[0].Country)

			// Now that we have Lat and Lon, let's query the weather
			units := "metric"
			if weatherAPIUnit == "F" {
				units = "imperial"
			}
			if hoursFromNow == 0 {
				url = "https://api.openweathermap.org/data/2.5/weather?lat=" + Lat + "&lon=" + Lon + "&units=" + units + "&appid=" + weatherAPIKey
			} else {
				url = "https://api.openweathermap.org/data/2.5/forecast?lat=" + Lat + "&lon=" + Lon + "&units=" + units + "&appid=" + weatherAPIKey
			}
			resp, err = http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, _ = io.ReadAll(resp.Body)
			weatherResponse := string(body)
			var openWeatherMapAPIResponse openWeatherMapAPIResponseStruct

			if hoursFromNow > 0 {
				// Forecast request: free API results are returned in 3 hours slots
				var openWeatherMapForecastAPIResponse openWeatherMapForecastAPIResponseStruct
				err = json.Unmarshal([]byte(weatherResponse), &openWeatherMapForecastAPIResponse)
				openWeatherMapAPIResponse = openWeatherMapForecastAPIResponse.List[hoursFromNow/3]
			} else {
				// Current weather request
				err = json.Unmarshal([]byte(weatherResponse), &openWeatherMapAPIResponse)
			}

			if err != nil {
				panic(err)
			}

			conditionCode := openWeatherMapAPIResponse.Weather[0].Id
			logger.Println(conditionCode)

			if conditionCode < 300 {
				// Thunderstorm
				condition = "Thunderstorms"
			} else if conditionCode < 400 {
				// Drizzle
				condition = "Rain"
			} else if conditionCode < 600 {
				// Rain
				condition = "Rain"
			} else if conditionCode < 700 {
				// Snow
				condition = "Snow"
			} else if conditionCode < 800 {
				// Athmosphere
				if openWeatherMapAPIResponse.Weather[0].Main == "Mist" ||
					openWeatherMapAPIResponse.Weather[0].Main == "Fog" {
					condition = "Rain"
				} else {
					condition = "Windy"
				}
			} else if conditionCode == 800 {
				// Clear
				if openWeatherMapAPIResponse.DT < openWeatherMapAPIResponse.Sys.Sunset {
					condition = "Sunny"
				} else {
					condition = "Stars"
				}
			} else if conditionCode < 900 {
				// Cloud
				condition = "Cloudy"
			} else {
				condition = openWeatherMapAPIResponse.Weather[0].Main
			}

			is_forecast = "false"
			t := time.Unix(int64(openWeatherMapAPIResponse.DT), 0)
			local_datetime = t.Format(time.RFC850)
			logger.Println(local_datetime)
			speakable_location_string = openWeatherMapAPIResponse.Name
			temperature = fmt.Sprintf("%f", math.Round(openWeatherMapAPIResponse.Main.Temp))
			if weatherAPIUnit == "C" {
				temperature_unit = "C"
			} else {
				temperature_unit = "F"
			}
		}
	} else {
		condition = "Snow"
		is_forecast = "false"
		local_datetime = "test"              // preferably local time in UTC ISO 8601 format ("2022-06-15 12:21:22.123")
		speakable_location_string = location // preferably the processed location
		temperature = "120"
		temperature_unit = "C"
	}
	return condition, is_forecast, local_datetime, speakable_location_string, temperature, temperature_unit
}
