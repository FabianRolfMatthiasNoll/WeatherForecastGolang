package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func displayWeather(firstIndex, LastIndex int, weather Weather, cityName string) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Time", "Temperature", "Rain", "Showers", "Cloud%", "Humidity"})

	fmt.Println("You are currently viewing the weather of: ", cityName)

	for i := firstIndex; i <= LastIndex; i = i + 2 {
		time := weather.Hourly.Time[i][11:13] + ":00"
		temp := weather.Hourly.Temperature2M[i]
		rain := weather.Hourly.Rain[i]
		shower := weather.Hourly.Showers[i]
		clouds := weather.Hourly.Cloudcover[i]
		humidity := weather.Hourly.Relativehumidity2M[i]
		//fmt.Print("|", weather.Hourly.Time[i][11:13], ":00")
		t.AppendRow([]interface{}{time, temp, rain, shower, clouds, humidity})
	}

	//fmt.Print("\nTemperature: ")
	for i := firstIndex; i <= LastIndex; i = i + 2 {
		//fmt.Print("|", weather.Hourly.Temperature2M[i])
	}
	//fmt.Print("\nCloudcoverage: ")
	for i := firstIndex; i <= LastIndex; i = i + 2 {
		//fmt.Print("|", weather.Hourly.Cloudcover[i])
	}
	//fmt.Print("\nHumidity: ")
	for i := firstIndex; i <= LastIndex; i = i + 2 {
		//fmt.Print("|", weather.Hourly.Relativehumidity2M[i])
	}
	//fmt.Print("\nRain: ")
	for i := firstIndex; i <= LastIndex; i = i + 2 {
		//fmt.Print("|", weather.Hourly.Rain[i])
	}
	//fmt.Print("\nShowers: ")
	for i := firstIndex; i <= LastIndex; i = i + 2 {
		//fmt.Print("|", weather.Hourly.Showers[i])
	}
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Render()
}
