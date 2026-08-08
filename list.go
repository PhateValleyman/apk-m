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
		fmt.Println(C("No database", RED))
		return
	}
	defer rows.Close()
	for rows.Next() {
		var (
			path    string
			pkg     string
			version string
			mod     int
		)
		rows.Scan(&path, &pkg, &version, &mod)
		fmt.Println("================================")
		fmt.Printf("      %s\n", path)
		fmt.Println("--------------------------------")
		fmt.Printf(" %s | %s\n", pkg, version)
		fmt.Println("================================")
	}
}
