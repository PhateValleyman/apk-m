package main

import (
	"fmt"
	"os"
)

func DatabaseExists() bool {
	_, err := os.Stat(
		DatabasePath(),
	)
	return err == nil
}

func FindCommand(path string) {
	Header("=== SEARCH & SCAN ===")
	if DatabaseExists() {
		Warn("Database already exists. Overwrite? [y/N]")
		var answer string
		fmt.Scanln(
			&answer,
		)
		if answer != "y" && answer != "Y" {
			Info("Operation cancelled by user.")
			return
		}
	}
	Find(path)
}

func ListCommand() {
	Header("=== APPLICATION LIST ===")
	if !DatabaseExists() {
		Error("No database found. Please run find first.")
		return
	}
	List()
}

func SortCommand() {
	Header("=== SORTING APPLICATIONS ===")
	if !DatabaseExists() {
		Error("No database found. Please run find first.")
		return
	}
	Sort()
}

func DuplicateCommand() {
	Header("=== DUPLICATE CHECK ===")
	if !DatabaseExists() {
		Error("No database found. Please run find first.")
		return
	}
	InitDB()
	RemoveDuplicates()
}

func CleanCommand() {
	Header("=== CLEANING DATABASE ===")
	ResetDB()
	Success("Database has been reset.")
}
