package main
import (
	"database/sql"
	"fmt"
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
	fmt.Println(
		C(
			"Database removed",
			YELLOW,
		),
	)
}
