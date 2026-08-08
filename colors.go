package main

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	CYAN   = "\033[36m"
	WHITE  = "\033[37m"
)

func C(
	text string,
	color string,
) string {
	return color + text + RESET
}
