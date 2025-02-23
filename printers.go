package main

import (
	"fmt"
	"strconv"

	"github.com/KartoffelChipss/gowttr/utils"
	"github.com/KartoffelChipss/gowttr/weather"
	"github.com/fatih/color"
)

var bold = color.New(color.Bold).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var italic = color.New(color.Italic).SprintFunc()
var darkGray = color.New(color.FgHiBlack).SprintFunc()
var blueBg = color.New(color.BgBlue).SprintFunc()

func PrintVersion() {
	fmt.Printf("%s\n", APP_VERSION)
}

func PrintHelp() {
	fmt.Printf("%s %s %s\n", bold("Usage:"), cyan("gowttr"), green("[flags]"))
	fmt.Println("")
	fmt.Println(bold("Flags:"))
	printFlags([]Flag{
		{"-l", "<location>", "Location to get the weather for"},
		{"-v", "", "Print the app version"},
		{"-h", "", "Print the help page"},
	})
	fmt.Println("")
	fmt.Println(bold("Examples:"))
	fmt.Println(darkGray("  gowttr -location Berlin"))
	fmt.Println(darkGray("  gowttr -location 'New York'"))
}

type Flag struct {
	Short       string
	Argument    string
	Description string
}

func printFlags(flags []Flag) {
	longestShortArgCombo := 0

	for _, flag := range flags {
		if len(flag.Short)+len(flag.Argument)+1 > longestShortArgCombo {
			longestShortArgCombo = len(flag.Short) + len(flag.Argument) + 1
		}
	}

	for _, flag := range flags {
		fmt.Printf(
			"  %s %s %s %s\n",
			green(flag.Short),
			yellow(flag.Argument),
			getSpaces(longestShortArgCombo-len(flag.Short)-len(flag.Argument)),
			italic(flag.Description),
		)
	}
}

func getSpaces(count int) string {
	spaces := ""
	for range count {
		spaces += " "
	}
	return spaces
}

func PrintWeather(weatherToPrint weather.WeatherData) {
	fmt.Printf(blueBg("%s, %s\n\n\033[0m"), weatherToPrint.NearestArea[0].AreaName[0].Value, weatherToPrint.NearestArea[0].Country[0].Value)

	weatherDesc := weatherToPrint.CurrentCondition[0].WeatherDesc[0].Value
	temp := formatTemp(weatherToPrint.CurrentCondition[0].TempC, weatherToPrint.CurrentCondition[0].FeelsLikeC)
	precip := utils.FormatText(fmt.Sprintf("&#83a5fa%v mm", weatherToPrint.CurrentCondition[0].PrecipMM))
	pressure := fmt.Sprintf("%v hPa", weatherToPrint.CurrentCondition[0].Pressure)
	wind := fmt.Sprintf("%s %v km/h", utils.GetWindDirPosArrow(weatherToPrint.CurrentCondition[0].WindDir16Point), weatherToPrint.CurrentCondition[0].WindSpeedKmph)

	fmt.Printf(weather.GetWeatherArt(weatherDesc)+"\n\n", weatherDesc, temp, wind, precip, pressure)
}

func formatTemp(temp string, feelsLikeTemp string) string {
	tmpInt, err := strconv.Atoi(temp)
	if err != nil {
		return utils.FormatText("&cInvalid temperature value: " + temp)
	}

	feelsLikeInt, err := strconv.Atoi(feelsLikeTemp)
	if err != nil {
		return utils.FormatText("&cInvalid feels-like temperature value: " + feelsLikeTemp)
	}

	var tempString string
	if tmpInt > 0 {
		tempString = utils.FormatText(fmt.Sprintf("&%s+%v", getTempColor(tmpInt), temp))
	} else {
		tempString = utils.FormatText(fmt.Sprintf("&%s%v", getTempColor(tmpInt), temp))
	}

	var feelsLikeTempString string
	if feelsLikeInt > 0 {
		feelsLikeTempString = fmt.Sprintf("+%v", feelsLikeTemp)
	} else {
		feelsLikeTempString = fmt.Sprintf("%v", feelsLikeTemp)
	}

	return fmt.Sprintf("%s°C (Feels like %s°C)", tempString, feelsLikeTempString)
}

func getTempColor(temp int) string {
	if temp < -10 {
		return "#4C6FBF"
	} else if temp < 0 {
		return "#7EA8F8"
	} else if temp < 10 {
		return "#A3C9FA"
	} else if temp < 20 {
		return "#7ACB98"
	} else if temp < 25 {
		return "#B0D96C"
	} else if temp < 30 {
		return "#F7DC81"
	} else if temp < 35 {
		return "#F4A977"
	} else if temp < 40 {
		return "#E27B7B"
	} else {
		return "#C96A6A"
	}
}
