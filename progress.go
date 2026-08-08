package main

import (
	"fmt"
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
	bar := ""
	for i := 0; i < width; i++ {
		if i < done {
			bar += "#"
		} else {
			bar += "-"
		}
	}
	fmt.Printf(
		"\r%s [%s] %d%%",
		text,
		bar,
		percent,
	)
	if current >= total {
		fmt.Println()
	}
}
