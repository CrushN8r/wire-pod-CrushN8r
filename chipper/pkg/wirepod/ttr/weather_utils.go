package wirepod_ttr

import (
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/kercre123/wire-pod/chipper/pkg/logger"
	lcztn "github.com/kercre123/wire-pod/chipper/pkg/wirepod/localization"
)

func removeEndPunctuation(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	lastIndex := len(runes) - 1

	if unicode.IsPunct(runes[lastIndex]) {
		return string(runes[:lastIndex])
	}

	return s
}

type WeatherStruct struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type openWeatherMapAPIResponseStruct struct {
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Weather []WeatherStruct `json:"weather"`
	Base    string          `json:"base"`
	Main    struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	DT  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type openWeatherMapForecastAPIResponseStruct struct {
	Cod     string                            `json:"cod"`
	Message int                               `json:"message"`
	Cnt     int                               `json:"cnt"`
	List    []openWeatherMapAPIResponseStruct `json:"list"`
}

func weatherParser(speechText string, botLocation string, botUnits string) (string, string, string, string, string, string) {
	var specificLocation bool
	var apiLocation string
	var speechLocation string
	var hoursFromNow int
	if strings.Contains(speechText, lcztn.GetText(lcztn.STR_WEATHER_IN)) {
		splitPhrase := strings.SplitAfter(removeEndPunctuation(speechText), lcztn.GetText(lcztn.STR_WEATHER_IN))
		speechLocation = strings.TrimSpace(splitPhrase[1])
		if os.Getenv("STT_SERVICE") != "whisper.cpp" {
			if len(splitPhrase) == 3 {
				speechLocation = speechLocation + " " + strings.TrimSpace(splitPhrase[2])
			} else if len(splitPhrase) == 4 {
				speechLocation = speechLocation + " " + strings.TrimSpace(splitPhrase[2]) + " " + strings.TrimSpace(splitPhrase[3])
			} else if len(splitPhrase) > 4 {
				speechLocation = speechLocation + " " + strings.TrimSpace(splitPhrase[2]) + " " + strings.TrimSpace(splitPhrase[3])
			}
			splitLocation := strings.Split(speechLocation, " ")
			if len(splitLocation) == 2 {
				speechLocation = splitLocation[0] + ", " + splitLocation[1]
			} else if len(splitLocation) == 3 {
				speechLocation = splitLocation[0] + " " + splitLocation[1] + ", " + splitLocation[2]
			}
		}
		logger.Println("Location parsed from speech: " + "`" + speechLocation + "`")
		specificLocation = true
	} else {
		logger.Println("No location parsed from speech")
		specificLocation = false
	}
	hoursFromNow = 0
	hours, _, _ := time.Now().Clock()
	if strings.Contains(speechText, lcztn.GetText(lcztn.STR_WEATHER_THIS_AFTERNOON)) {
		if hours < 14 {
			hoursFromNow = 14 - hours
		}
	} else if strings.Contains(speechText, lcztn.GetText(lcztn.STR_WEATHER_TONIGHT)) {
		if hours < 20 {
			hoursFromNow = 20 - hours
		}
	} else if strings.Contains(speechText, lcztn.GetText(lcztn.STR_WEATHER_THE_DAY_AFTER_TOMORROW)) {
		hoursFromNow = 24 - hours + 24 + 9
	} else if strings.Contains(speechText, lcztn.GetText(lcztn.STR_WEATHER_FORECAST)) ||
		strings.Contains(speechText, lcztn.GetText(lcztn.STR_WEATHER_TOMORROW)) {
		hoursFromNow = 24 - hours + 9
	}
	logger.Println("Looking for forecast " + strconv.Itoa(hoursFromNow) + " hours from now...")

	if specificLocation {
		apiLocation = speechLocation
	} else {
		apiLocation = botLocation
	}
	// call to weather API
	condition, is_forecast, local_datetime, speakable_location_string, temperature, temperature_unit := getWeather(apiLocation, botUnits, hoursFromNow)
	return condition, is_forecast, local_datetime, speakable_location_string, temperature, temperature_unit
}
