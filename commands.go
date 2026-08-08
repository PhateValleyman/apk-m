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
	if DatabaseExists() {
		fmt.Println(
			C(
				"Database already exists. Overwrite? [y/N]",
				YELLOW,
			),
		)
		var answer string
		fmt.Scanln(
			&answer,
		)
		if answer != "y" && answer != "Y" {
			return
		}
	}
	Find(path)
}
func ListCommand() {
	if !DatabaseExists() {
		fmt.Println(
			C(
				"No database found. Run find first.",
				RED,
			),
		)
		return
	}
	List()
}
func SortCommand() {
	if !DatabaseExists() {
		fmt.Println(
			C(
				"No database found. Run find first.",
				RED,
			),
		)
		return
	}
	Sort()
}
func DuplicateCommand() {
	if !DatabaseExists() {
		fmt.Println(
			C(
				"No database found. Run find first.",
				RED,
			),
		)
		return
	}
	InitDB()
	RemoveDuplicates()
}
func CleanCommand() {
	ResetDB()
}
