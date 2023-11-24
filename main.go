package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC      float64 `json:"temp_c"`
		FeelsLikeC float64 `json:"feelslike_c"`
		Condition  struct {
			Text string `json:"text"`
		}
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			MaxTempC  float64 `json:"maxtemp_c"`
			MinTempC  float64 `json:"mintemp_c"`
			AvgTempC  float64 `json:"avgtemp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
			Hour []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TempC        float64 `json:"temp_c"`
				ChanceOfRain int     `json:"chance_of_rain"`
				ChanceOfSnow int     `json:"chance_of_snow"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type ConnectionError struct {
	statusCode int32
	msg        string
}

func (c *ConnectionError) Error() string {
	return fmt.Sprintf("status: %d, msg: %s", c.statusCode, c.msg)
}

func main() {
	a := "6b04271bb8824900be9225552232411"
	q := "Bytom"

	if len(os.Args) > 2 {
		q = os.Args[4]
	}

	url := "https://api.weatherapi.com/v1/forecast.json?key=" + a + "&q=" + q + "&days=1&alerts=yes"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
		/* if resp.StatusCode != 200 {
			error := ConnectionError{
				statusCode: int32(resp.StatusCode),
				msg:        "WeatherAPI is not available. Try again later.",
			}
			color.Red(error.Error())
			return
		} */
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(body))
}
