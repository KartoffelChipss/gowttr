package weather

import (
	"strings"

	"github.com/KartoffelChipss/gowttr/utils"
)

func GetWeatherArt(weatherDesc string) string {
	weatherDesc = strings.ToLower(weatherDesc)

	switch weatherDesc {
	case "sunny", "clear":
		return utils.FormatText("&#fffc3e     \\   /     &r%s\n&#fffc3e      .-.      &r%s\n&#fffc3e   ― (   ) ―   &r%s\n&#fffc3e      `-’      &r%s\n&#fffc3e     /   \\     &r%s")
	case "partly cloudy":
		return utils.FormatText("&#fffc3e    \\  /       &r%s\n  &#fffc3e_ /\"\"&#b3b3b3.-.     &r%s\n&#fffc3e    \\_&#b3b3b3(   ).   &r%s\n    &#fffc3e/&#b3b3b3(___(__)  &r%s\n               &r%s")
	case "overcast", "shower in vicinity":
		return utils.FormatText("               &r%s\n&#b3b3b3      .--.     &r%s\n&#b3b3b3   .-(    ).   &r%s\n&#b3b3b3  (___.__)__)  &r%s\n               &r%s")
	case "patchy rain", "patchy rain nearby":
		return utils.FormatText("&#fffc3e  _`/\"\"&#b3b3b3.-.     &r%s\n   &#fffc3e,\\_&#b3b3b3(   ).   &r%s\n    &#fffc3e/&#b3b3b3(___(__)  &r%s\n      &#83a5fa‘ ‘ ‘ ‘  &r%s\n     &#83a5fa‘ ‘ ‘ ‘   &r%s")
	case "light drizzle", "drizzle":
		return utils.FormatText("      &#b3b3b3.-.      &r%s\n     &#b3b3b3(   ).    &r%s\n    &#b3b3b3(___(__)   &r%s\n     &#83a5fa‘ ‘ ‘ ‘   &r%s\n    &#83a5fa‘ ‘ ‘ ‘    &r%s")
	case "mist", "freezing fog":
		return utils.FormatText("               &r%s\n&#b3b3b3  _ - _ - _ -  &r%s\n&#b3b3b3   _ - _ - _   &r%s\n&#b3b3b3  _ - _ - _ -  &r%s\n               &r%s")
	case "light snow":
		return utils.FormatText("  &#fffc3e_`/\"\"&#b3b3b3.-.     &r%s\n   &#fffc3e,\\_&#b3b3b3(   ).   &r%s\n    &#fffc3e/&#b3b3b3(___(__)  &r%s\n      &#ffffff*  *  *  &r%s\n     &#ffffff*  *  *   &r%s")
	case "moderate snow", "heavy snow":
		return utils.FormatText("      &#b3b3b3.-.      &r%s\n     &#b3b3b3(   ).    &r%s\n    &#b3b3b3(___(__)   &r%s\n    &#ffffff* * * *    &r%s\n   &#ffffff* * * *     &r%s")
	case "light freezing", "light sleet showers":
		return utils.FormatText("      &#b3b3b3.-.      &r%s\n     &#b3b3b3(   ).    &r%s\n    &#b3b3b3(___(__)   &r%s\n     &#83a5fa‘ &#ffffff* &#83a5fa‘ &#ffffff*   &r%s\n    &#ffffff* &#83a5fa‘ &#ffffff* &#83a5fa‘    &r%s")
	default:
		return utils.FormatText("&r%s\n&r%s\n&r%s\n&r%s\n&r%s")
	}
}
