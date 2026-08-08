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
	SetTitle(
		"searching " + path,
	)
	fmt.Println(
		C(
			"[SEARCH]",
			CYAN,
		),
		path,
	)
	files := make(chan string, 256)
	var wg sync.WaitGroup
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
			}
		}()
	}
	err := filepath.WalkDir(
		path,
		func(
			current string,
			entry os.DirEntry,
			err error,
		) error {
			if err != nil {
				return nil
			}
			if entry.IsDir() {
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
				files <- current
			}
			return nil
		},
	)
	if err != nil {
		fmt.Println(err)
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
