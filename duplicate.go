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
		Error("Failed to check for duplicates: %v", err)
		return
	}
	defer rows.Close()
	
	count := 0
	for rows.Next() {
		count++
		var (
			hash  string
			num   int
		)
		rows.Scan(
			&hash,
			&num,
		)
		Warn("Duplicate found: %s (%s copies)", C(hash, BCYAN), C(fmt.Sprintf("%d", num), BYELLOW+BOLD))
	}
	
	if count > 0 {
		Info("Found %d unique hashes with duplicates.", count)
	} else {
		Success("No duplicate applications found.")
	}
}
