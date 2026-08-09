package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Find(path string) {
	cfg := LoadConfig()
	if path == "" {
		fmt.Print(
			"Search path: ",
		)
		fmt.Scanln(
			&path,
		)
	}
	cfg.StartPath = path
	cfg.ArchivePath = filepath.Join(
		path,
		"SORTED-APPS",
	)
	SaveConfig(cfg)
	InitDB()
	ResetDB()
	InitDB()
	fmt.Println(
		C(
			"[SEARCH]",
			CYAN,
		),
		path,
	)
	var allFiles []string
	filepath.WalkDir(
		path,
		func(
			current string,
			entry os.DirEntry,
			err error,
		) error {
			if err != nil || entry.IsDir() {
				return nil
			}
			ext := strings.ToLower(
				filepath.Ext(
					current,
				),
			)
			switch ext {
			case ".apk",
				".apks",
				".xapk",
				".apkm":
				allFiles = append(allFiles, current)
			}
			return nil
		},
	)
	total := len(allFiles)
	if total == 0 {
		fmt.Println(C("No APK files found.", YELLOW))
		return
	}
	SetTitle(
		fmt.Sprintf("scanning %d files in %s", total, path),
	)
	files := make(chan string, 256)
	var wg sync.WaitGroup
	var mu sync.Mutex
	currentCount := 0
	workers := cfg.Workers
	if workers < 1 {
		workers = 4
	}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range files {
				ScanAPK(
					file,
				)
				mu.Lock()
				currentCount++
				Progress(currentCount, total, "Scanning")
				mu.Unlock()
			}
		}()
	}
	for _, file := range allFiles {
		files <- file
	}
	close(files)
	wg.Wait()
	ClearTitle()
	Bell()
	fmt.Println()
	fmt.Println(
		C(
			"Search finished",
			GREEN,
		),
	)
}
func SaveRecord(
	path string,
	start string,
	typ string,
	name string,
	pkg string,
	version string,
	code string,
	hash string,
	size int64,
	signature string,
	status string,
	isMod int,
) {
	_, err := DB.Exec(`
INSERT INTO apps
(
path,
filename,
start_path,
type,
name,
package_id,
version_name,
version_code,
sha256,
size,
signature,
signature_status,
is_mod,
created
)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?,?,strftime('%s','now'))
`,
		path,
		filepath.Base(path),
		start,
		typ,
		name,
		pkg,
		version,
		code,
		hash,
		size,
		signature,
		status,
		isMod,
	)
	if err != nil {
		fmt.Println(err)
	}
}
