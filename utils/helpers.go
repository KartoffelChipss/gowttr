package utils

func GetWindDirPosArrow(windDir16Point string) string {
	directions := map[string]string{
		"N": "↑", "NNE": "↗", "NE": "→", "ENE": "↘",
		"E": "↓", "ESE": "↙", "SE": "←", "SSE": "↖",
		"S": "↑", "SSW": "↗", "SW": "→", "WSW": "↘",
		"W": "↓", "WNW": "↙", "NW": "←", "NNW": "↖",
	}
	if arrow, exists := directions[windDir16Point]; exists {
		return arrow
	}
	return " "
}
