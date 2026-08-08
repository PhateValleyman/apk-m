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
		centerPrint(pkg+" | "+version, width)
		fmt.Println(line)
	}
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
