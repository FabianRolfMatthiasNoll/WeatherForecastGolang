package main

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"strings"
	// https://github.com/jedib0t/go-pretty Repository of the pretty tables
)

func getCity() (city string) {
	fmt.Print("Please enter the city you want a weather report from: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	city = strings.TrimRight(input, "\r\n")
	return
}

func displayWeather(firstIndex, LastIndex int, weather Weather, cityName string) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Time", "Temperature", "Rain", "Showers", "Clouds", "Humidity"})

	fmt.Println("You are currently viewing the weather of: ", cityName)

	for i := firstIndex; i <= LastIndex; i = i + 2 {
		time := weather.Hourly.Time[i][11:13] + ":00"
		temp := fmt.Sprintf("%.1f°C", weather.Hourly.Temperature2M[i])
		rain := fmt.Sprintf("%.1fmm", weather.Hourly.Rain[i])
		shower := fmt.Sprintf("%dmm", weather.Hourly.Showers[i])
		clouds := fmt.Sprintf("%d", weather.Hourly.Cloudcover[i]) + "%"
		humidity := fmt.Sprintf("%d", weather.Hourly.Relativehumidity2M[i]) + "%"
		//fmt.Print("|", weather.Hourly.Time[i][11:13], ":00")
		t.AppendRow([]interface{}{time, temp, rain, shower, clouds, humidity})
	}
	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:   "Temperature",
			Align:  text.AlignRight,
			VAlign: text.VAlignMiddle,
		},
		{
			Name:   "Rain",
			Align:  text.AlignRight,
			VAlign: text.VAlignMiddle,
		},
		{
			Name:   "Showers",
			Align:  text.AlignRight,
			VAlign: text.VAlignMiddle,
		},
		{
			Name:   "Clouds",
			Align:  text.AlignRight,
			VAlign: text.VAlignMiddle,
		},
		{
			Name:   "Humidity",
			Align:  text.AlignRight,
			VAlign: text.VAlignMiddle,
		},
	})
	t.Render()
	//Placeholder so that the .exe stays open after displaying the weather
	fmt.Scanf("h")
}
