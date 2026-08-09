package main

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

var DB *sql.DB

func DatabasePath() string {
	return filepath.Join(
		os.Getenv("HOME"),
		".config",
		"apk-m",
		"list.db",
	)
}
func InitDB() {
	var err error
	DB, err = sql.Open(
		"sqlite",
		DatabasePath(),
	)
	if err != nil {
		panic(err)
	}
	query := `
CREATE TABLE IF NOT EXISTS apps (
	id INTEGER PRIMARY KEY,
	path TEXT,
	filename TEXT,
	start_path TEXT,
	type TEXT,
	name TEXT,
	package_id TEXT,
	version_name TEXT,
	version_code TEXT,
	sha256 TEXT,
	size INTEGER,
	signature TEXT,
	signature_status TEXT,
	is_mod INTEGER DEFAULT 0,
	created INTEGER
);
`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
}
func ResetDB() {
	os.Remove(
		DatabasePath(),
	)
}

func CreateIndexes() {

	DB.Exec("CREATE INDEX IF NOT EXISTS idx_package ON apps(package_id)")

	DB.Exec("CREATE INDEX IF NOT EXISTS idx_hash ON apps(sha256)")

	DB.Exec("CREATE INDEX IF NOT EXISTS idx_mod ON apps(is_mod)")

}
