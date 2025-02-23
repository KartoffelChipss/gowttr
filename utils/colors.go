package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FormatText(text string) string {
	var colors = map[string]string{
		"&r": "\033[0m",

		// Text Colors
		"&0": "\033[30m",
		"&4": "\033[31m",
		"&2": "\033[32m",
		"&6": "\033[33m",
		"&1": "\033[34m",
		"&5": "\033[35m",
		"&3": "\033[36m",
		"&7": "\033[37m",

		// Bright Text Colors
		"&8": "\033[90m",
		"&c": "\033[91m",
		"&a": "\033[92m",
		"&e": "\033[93m",
		"&9": "\033[94m",
		"&d": "\033[95m",
		"&b": "\033[96m",
		"&f": "\033[97m",

		// Formatting
		"&l": "\033[1m",
		"&p": "\033[2m", //dim
		"&o": "\033[3m",
		"&n": "\033[4m",
		"&q": "\033[5m", //blink
		"&s": "\033[7m", //reverse
		"&t": "\033[8m", //hidden
		"&m": "\033[9m", //strikethrough
	}

	text = text + colors["&r"]

	for k, v := range colors {
		text = strings.ReplaceAll(text, k, v)
	}

	hexPattern := regexp.MustCompile(`&#([0-9a-fA-F]{6})`)

	text = hexPattern.ReplaceAllStringFunc(text, func(match string) string {
		hex := match[2:]
		r, _ := strconv.ParseInt(hex[0:2], 16, 64)
		g, _ := strconv.ParseInt(hex[2:4], 16, 64)
		b, _ := strconv.ParseInt(hex[4:6], 16, 64)

		return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	})

	return text
}
