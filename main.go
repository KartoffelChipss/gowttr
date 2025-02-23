package main

import (
	"os"

	"github.com/KartoffelChipss/gowttr/weather"
	"github.com/fatih/color"
)

func main() {
	var location string

	for i, arg := range os.Args {
		if arg == "-l" && i+1 < len(os.Args) {
			location = os.Args[i+1]
		}

		if arg == "-v" {
			PrintVersion()
			return
		}

		if (arg == "-h" || arg == "--help") && i == 1 {
			PrintHelp()
			return
		}
	}

	weather, err := weather.GetWeather(location)

	if err != nil {
		color.Red("Error: %s", err)
		return
	}

	PrintWeather(weather)
}
