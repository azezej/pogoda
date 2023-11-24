package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
)

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
		q = os.Args[1]
	}

	url := "https://api.weatherapi.com/v1/forecast.json?key=" + a + "&q=" + q + "&days=1&aqi=yes&alerts=yes"
	resp, err := http.Get(url)
	if err != nil {
		if resp.StatusCode != 200 {
			error := ConnectionError{
				statusCode: int32(resp.StatusCode),
				msg:        "WeatherAPI is not available. Try again later.",
			}
			color.Red(error.Error())
			return
		}
		panic(err)
	}

	defer resp.Body.Close()
}
