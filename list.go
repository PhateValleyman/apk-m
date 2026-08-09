package main

import (
	"fmt"
	"strings"
)

func List() {
	InitDB()
	rows, err := DB.Query(`SELECT path, package_id, version_name, is_mod FROM apps ORDER BY package_id`)
	if err != nil {
		Error("Failed to query database: %v", err)
		return
	}
	defer rows.Close()

	width := GetTerminalWidth()
	if width < 60 {
		width = 60
	}

	count := 0
	for rows.Next() {
		count++
		var (
			path    string
			pkg     string
			version string
			mod     int
		)
		rows.Scan(&path, &pkg, &version, &mod)

		modTag := ""
		if mod == 1 {
			modTag = C(" [MOD]", BRED+BOLD)
		}

		fmt.Println(C(strings.Repeat("━", width), BMAGENTA))
		fmt.Printf("%s %s%s\n", C("Package:", BYELLOW), C(pkg, BWHITE+BOLD), modTag)
		fmt.Printf("%s %s\n", C("Version:", BYELLOW), C(version, BCYAN))
		fmt.Printf("%s %s\n", C("Path:   ", BYELLOW), C(path, DIM))
	}
	fmt.Println(C(strings.Repeat("━", width), BMAGENTA))
	Success("Total applications found: %d", count)
}
func centerString(text string, width int) string {
	if len(text) > width-2 {
		text = text[:width-5] + "..."
	}
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	leftPad := strings.Repeat(" ", padding)
	rightPad := strings.Repeat(" ", width-len(text)-padding)
	return leftPad + text + rightPad
}
func centerPrint(text string, width int) {
	if len(text) > width-4 {
		text = "..." + text[len(text)-(width-7):]
	}
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Printf("%s%s\n", strings.Repeat(" ", padding), text)
}
