package main
import (
	"fmt"
	"strings"
)
func List() {
	InitDB()
	rows, err := DB.Query(`SELECT path, package_id, version_name, is_mod FROM apps ORDER BY package_id`)
	if err != nil {
		fmt.Println(C("No database", RED))
		return
	}
	defer rows.Close()
	width := GetTerminalWidth()
	if width < 40 {
		width = 40
	}
	for rows.Next() {
		var (
			path    string
			pkg     string
			version string
			mod     int
		)
		rows.Scan(&path, &pkg, &version, &mod)
		line := strings.Repeat("=", width)
		sep := strings.Repeat("-", width)
		fmt.Println(line)
		centerPrint(path, width)
		fmt.Println(sep)
		halfWidth := width / 2
		left := centerString(pkg, halfWidth)
		right := centerString(version, width-halfWidth-1)
		fmt.Printf("%s|%s\n", left, right)
		fmt.Println(line)
	}
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
