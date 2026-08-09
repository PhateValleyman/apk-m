package main

import (
	"fmt"
	"strings"
)

func Progress(
	current int,
	total int,
	text string,
) {
	width := 30
	percent := 0
	if total > 0 {
		percent = current * 100 / total
	}
	done := width * percent / 100
	bar := C(strings.Repeat("━", done), BGREEN) + C(strings.Repeat("─", width-done), DIM)
	
	fmt.Printf(
		"\r%s [%s] %s",
		text,
		bar,
		C(fmt.Sprintf("%3d%%", percent), BYELLOW+BOLD),
	)
	if current >= total {
		fmt.Println()
	}
}
