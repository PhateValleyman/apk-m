package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "2.0"

func main() {
	var (
		showVersion bool
		showHelp    bool
		findPath    string
		list        bool
		sort        bool
		duplicate   bool
		clean       bool
	)

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.StringVar(&findPath, "f", "", "scan APK/APKS/XAPK/APKM in <path>")
	flag.StringVar(&findPath, "find", "", "scan APK/APKS/XAPK/APKM in <path>")
	flag.BoolVar(&list, "l", false, "list all applications from database")
	flag.BoolVar(&list, "list", false, "list all applications from database")
	flag.BoolVar(&sort, "s", false, "sort applications into folders")
	flag.BoolVar(&sort, "sort", false, "sort applications into folders")
	flag.BoolVar(&duplicate, "d", false, "show duplicate applications")
	flag.BoolVar(&duplicate, "duplicates", false, "show duplicate applications")
	flag.BoolVar(&clean, "c", false, "clean database")
	flag.BoolVar(&clean, "clean", false, "clean database")

	flag.Usage = ShowHelp
	flag.Parse()

	if showVersion {
		fmt.Printf("%s v%s\n", C("APK Manager", BCYAN+BOLD), C(VERSION, BYELLOW))
		return
	}
	if showHelp {
		ShowHelp()
		return
	}
	if findPath != "" {
		FindCommand(findPath)
		return
	}
	if list {
		ListCommand()
		return
	}
	if sort {
		SortCommand()
		return
	}
	if duplicate {
		DuplicateCommand()
		return
	}
	if clean {
		CleanCommand()
		return
	}

	if len(os.Args) == 1 {
		ShowHelp()
		return
	}
}

func ShowHelp() {
	fmt.Printf("%s v%s\n\n", C("APK Manager", BCYAN+BOLD), C(VERSION, BYELLOW))
	fmt.Println(C("Usage:", BYELLOW+BOLD))
	fmt.Println("  apk-m [OPTIONS]")
	fmt.Println()
	fmt.Println(C("Options:", BYELLOW+BOLD))
	fmt.Println("  -f, --find <PATH>      Scan APK/APKS/XAPK/APKM")
	fmt.Println("  -l, --list             Show database")
	fmt.Println("  -s, --sort             Sort applications")
	fmt.Println("  -d, --duplicates       Show duplicate applications")
	fmt.Println("  -c, --clean            Clean database")
	fmt.Println("  -v, --version          Show version")
	fmt.Println("  -h, --help             Show this help")
	fmt.Println()
	fmt.Println(C("Examples:", BYELLOW+BOLD))
	fmt.Println("  apk-m -f /storage/65D9-1787")
	fmt.Println("  apk-m --list")
	fmt.Println("  apk-m -s")
	fmt.Println()
}
