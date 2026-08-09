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
		catColor := BGREEN
		if isMod == 1 {
			category = "MOD"
			catColor = BRED
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
		
		fmt.Printf("\n%s %s\n", C("MOVE:", catColor+BOLD), C(filepath.Base(path), BWHITE))
		fmt.Printf("  %s %s\n", C("From:", BYELLOW), C(path, DIM))
		fmt.Printf("  %s   %s\n", C("To:", BYELLOW), C(dst, DIM))

		err := os.Rename(
			path,
			dst,
		)
		if err != nil {
			Error("Failed to move file: %v", err)
		}
		current++
		Progress(
			current,
			total,
			C("Sorting", BCYAN),
		)
	}
	ClearTitle()
	Bell()
	fmt.Println()
	Success("Sorting finished!")
	fmt.Printf("%s %s\n", C("Base Path:", BYELLOW), C(cfg.StartPath, BBLUE))
	fmt.Printf("%s %s\n", C("Original: ", BYELLOW), C(fmt.Sprintf("%d", original), BGREEN))
	fmt.Printf("%s %s\n", C("Mods:     ", BYELLOW), C(fmt.Sprintf("%d", mods), BRED))
	fmt.Printf("%s %s\n", C("Total Size:", BYELLOW), C(fmt.Sprintf("%.2f GB", float64(size)/1024/1024/1024), BCYAN))
	fmt.Println()
	RemoveDuplicates()
}
