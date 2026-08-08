package main
import (
	"fmt"
)
func RemoveDuplicates() {
	rows, err := DB.Query(`
SELECT
sha256,
COUNT(*)
FROM apps
GROUP BY sha256
HAVING COUNT(*) > 1
`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var (
			hash  string
			count int
		)
		rows.Scan(
			&hash,
			&count,
		)
		fmt.Println(
			C(
				"DUPLICATE:",
				YELLOW,
			),
			hash,
			count,
		)
	}
}
