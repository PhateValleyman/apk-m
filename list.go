package main

import (
	"fmt"
	"strings"
)

func List() {
	InitDB()
	rows, err := DB.Query(`
SELECT
path,
package_id,
version_name,
is_mod
FROM apps
ORDER BY package_id
`)
	if err != nil {
		fmt.Println(
			C(
				"No database",
				RED,
			),
		)
		return
	}
	defer rows.Close()

	fmt.Println()
	fmt.Printf(
		"%-50s | %-30s | %-15s | %-10s\n",
		C("PATH", CYAN),
		C("PACKAGE", CYAN),
		C("VERSION", CYAN),
		C("TYPE", CYAN),
	)
	fmt.Println(strings.Repeat("-", 115))

	for rows.Next() {
		var (
			path    string
			pkg     string
			version string
			mod     int
		)
		rows.Scan(
			&path,
			&pkg,
			&version,
			&mod,
		)
		typeName := "ORIGINAL"
		color := GREEN
		if mod == 1 {
			typeName = "MOD"
			color = RED
		}

		// Truncate path if it's too long
		displayPath := path
		if len(displayPath) > 50 {
			displayPath = "..." + displayPath[len(displayPath)-47:]
		}

		fmt.Printf(
			"%-50s | %-30s | %-15s | %-10s\n",
			displayPath,
			pkg,
			version,
			C(
				typeName,
				color,
			),
		)
	}
	fmt.Println()
}
