package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Sort() {
	cfg := LoadConfig()
	InitDB()
	SetTitle(
		"sorting " + cfg.StartPath,
	)
	rows, err := DB.Query(`
SELECT
path,
type,
package_id,
version_name,
version_code,
is_mod,
size
FROM apps
ORDER BY size DESC
`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	total := 0
	DB.QueryRow(
		"SELECT COUNT(*) FROM apps",
	).Scan(
		&total,
	)
	current := 0
	original := 0
	mods := 0
	var size int64
	for rows.Next() {
		var (
			path     string
			typ      string
			pkg      string
			version  string
			code     string
			isMod    int
			fileSize int64
		)
		rows.Scan(
			&path,
			&typ,
			&pkg,
			&version,
			&code,
			&isMod,
			&fileSize,
		)
		size += fileSize
		category := "ORIGINAL"
		if isMod == 1 {
			category = "MOD"
			mods++
		} else {
			original++
		}
		dir := filepath.Join(
			cfg.ArchivePath,
			category,
			typ,
			pkg,
			version+"-"+code,
		)
		os.MkdirAll(
			dir,
			0755,
		)
		dst := filepath.Join(
			dir,
			filepath.Base(path),
		)
		fmt.Println()
		fmt.Println(
			C(
				"MOVE:",
				GREEN,
			),
			path,
		)
		fmt.Println(
			" -> ",
			dst,
		)
		err := os.Rename(
			path,
			dst,
		)
		if err != nil {
			fmt.Println(
				C(
					err.Error(),
					RED,
				),
			)
		}
		current++
		Progress(
			current,
			total,
			"Sorting",
		)
	}
	ClearTitle()
	Bell()
	fmt.Println()
	fmt.Println(
		C(
			"SORT FINISHED",
			GREEN,
		),
	)
	fmt.Println(
		"Path:",
		cfg.StartPath,
	)
	fmt.Println(
		"Original:",
		original,
	)
	fmt.Println(
		"Mods:",
		mods,
	)
	fmt.Printf(
		"Size: %.2f GB\n",
		float64(size)/1024/1024/1024,
	)
	fmt.Println()
	RemoveDuplicates()
}
