package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "2.0"

func main() {
	version := flag.Bool(
		"v",
		false,
		"show version",
	)
	versionLong := flag.Bool(
		"version",
		false,
		"show version",
	)
	help := flag.Bool(
		"h",
		false,
		"show help",
	)
	helpLong := flag.Bool(
		"help",
		false,
		"show help",
	)
	find := flag.String(
		"f",
		"",
		"find path",
	)
	findLong := flag.String(
		"find",
		"",
		"find path",
	)
	list := flag.Bool(
		"l",
		false,
		"list",
	)
	listLong := flag.Bool(
		"list",
		false,
		"list",
	)
	sort := flag.Bool(
		"s",
		false,
		"sort",
	)
	sortLong := flag.Bool(
		"sort",
		false,
		"sort",
	)
	duplicate := flag.Bool(
		"d",
		false,
		"show duplicates",
	)
	duplicateLong := flag.Bool(
		"duplicates",
		false,
		"show duplicates",
	)
	clean := flag.Bool(
		"c",
		false,
		"clean database",
	)
	cleanLong := flag.Bool(
		"clean",
		false,
		"clean database",
	)
	flag.Parse()
	if *version || *versionLong {
		fmt.Println(
			C(
				"APK Manager v"+VERSION,
				CYAN,
			),
		)
		return
	}
	if *help || *helpLong {
		ShowHelp()
		return
	}
	if *find != "" || *findLong != "" {
		path := *find
		if path == "" {
			path = *findLong
		}
		FindCommand(path)
		return
	}
	if *list || *listLong {
		ListCommand()
		return
	}
	if *sort || *sortLong {
		SortCommand()
		return
	}
	if *duplicate || *duplicateLong {
		DuplicateCommand()
		return
	}
	if *clean || *cleanLong {
		CleanCommand()
		return
	}
	if len(os.Args) == 1 {
		ShowHelp()
		return
	}
}
func ShowHelp() {
	fmt.Println()
	fmt.Println(
		C(
			"APK Manager v"+VERSION,
			CYAN,
		),
	)
	fmt.Println()
	fmt.Println(
		C(
			"Usage:",
			YELLOW,
		),
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -h | --help",
	)
	fmt.Println(
		"      Show help",
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -v | --version",
	)
	fmt.Println(
		"      Show version",
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -f | --find <PATH>",
	)
	fmt.Println(
		"      Scan APK/APKS/XAPK/APKM",
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -l | --list",
	)
	fmt.Println(
		"      Show database",
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -s | --sort",
	)
	fmt.Println(
		"      Sort applications",
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -d | --duplicates",
	)
	fmt.Println(
		"      Show duplicate applications",
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -c | --clean",
	)
	fmt.Println(
		"      Clean database",
	)
	fmt.Println()
	fmt.Println(
		C(
			"Examples:",
			GREEN,
		),
	)
	fmt.Println()
	fmt.Println(
		"  apk-m -f /storage/65D9-1787",
	)
	fmt.Println(
		"  apk-m -l",
	)
	fmt.Println(
		"  apk-m -s",
	)
	fmt.Println(
		"  apk-m -d",
	)
	fmt.Println()
}
