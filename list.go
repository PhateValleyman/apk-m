package main

import (
	"fmt"
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
	fmt.Println(
		C(
			"PATH | PACKAGE | VERSION | TYPE",
			CYAN,
		),
	)
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
		fmt.Println(
			path,
			"|",
			pkg,
			"|",
			version,
			"|",
			C(
				typeName,
				color,
			),
		)
	}
}
