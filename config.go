package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	StartPath        string `json:"start_path"`
	ArchivePath      string `json:"archive_path"`
	Workers          int    `json:"workers"`
	RemoveDuplicates bool   `json:"remove_duplicates"`
}

var ConfigPath = filepath.Join(
	os.Getenv("HOME"),
	".config",
	"apk-m",
	"config.json",
)

func DefaultConfig() Config {
	return Config{
		StartPath:        "",
		ArchivePath:      "",
		Workers:          8,
		RemoveDuplicates: true,
	}
}
func LoadConfig() Config {
	cfg := DefaultConfig()
	file, err := os.Open(
		ConfigPath,
	)
	if err != nil {
		return cfg
	}
	defer file.Close()
	json.NewDecoder(
		file,
	).Decode(
		&cfg,
	)
	return cfg
}
func SaveConfig(
	cfg Config,
) {
	os.MkdirAll(
		filepath.Dir(
			ConfigPath,
		),
		0755,
	)
	file, _ := os.Create(
		ConfigPath,
	)
	defer file.Close()
	json.NewEncoder(
		file,
	).Encode(
		cfg,
	)
}
