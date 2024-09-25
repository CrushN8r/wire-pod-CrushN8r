package wirepod_ttr

/* TODO:
Create seperate functions for weatherAPI and openweathermap,
create a standard for how weather functions should be created
*/

// *** WEATHERAPI.COM ***

type weatherAPIResponseStruct struct {
	Location struct {
		Name      string `json:"name"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
	} `json:"current"`
}
type weatherAPICladStruct []struct {
	APIValue string `json:"APIValue"`
	CladType string `json:"CladType"`
}

// *** OPENWEATHERMAP.ORG ***

type openWeatherMapAPIGeoCodingStruct struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state"`
}

/*
//3.0 API, requires your credit card even to get 1k free requests per day

type openWeatherMapAPIResponseStruct struct {
    Lat      			float64 `json:"lat"`
	Lon					float64 `json:"lon"`
	timezone		  	string `json:"timezone"`
	timezone_offset	  	string `json:"timezone_offset"`
	Current struct {
		DT	 	 	int     `json:"dt"`
		Sunrise	 	int     `json:"sunrise"`
		Sunset	 	int     `json:"sunset"`
		Temp	    float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Pressure	int     `json:"pressure"`
		Humidity	int     `json:"humidity"`
		DewPoint	float64 `json:"dew_point"`
		UVI	        float64 `json:"uvi"`
		Clouds	 	int     `json:"clouds"`
		Visibility	int     `json:"visibility"`
		WindSpeed	float64 `json:"wind_speed"`
		WindDeg	 	int     `json:"wid_deg"`
		WindGust	float64 `json:"wind_gust"`
		Weather        struct {
			Id	 		int    `json:"id"`
			Main 		string `json:"main"`
			Description string `json:"description"`
			Icon 		string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
}
*/

//2.5 API