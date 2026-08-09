package main

import "fmt"

const (
	RESET  = "\033[0m"
	BOLD   = "\033[1m"
	DIM    = "\033[2m"
	ITALIC = "\033[3m"
	UNDER  = "\033[4m"

	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	CYAN    = "\033[36m"
	WHITE   = "\033[37m"

	BRED     = "\033[91m"
	BGREEN   = "\033[92m"
	BYELLOW  = "\033[93m"
	BBLUE    = "\033[94m"
	BMAGENTA = "\033[95m"
	BCYAN    = "\033[96m"
	BWHITE   = "\033[97m"
)

func C(text string, color string) string {
	return color + text + RESET
}

func Info(format string, a ...interface{}) {
	fmt.Printf(C("[*] ", BCYAN)+format+"\n", a...)
}

func Success(format string, a ...interface{}) {
	fmt.Printf(C("[+] ", BGREEN)+format+"\n", a...)
}

func Warn(format string, a ...interface{}) {
	fmt.Printf(C("[!] ", BYELLOW)+format+"\n", a...)
}

func Error(format string, a ...interface{}) {
	fmt.Printf(C("[-] ", BRED)+format+"\n", a...)
}

func Header(text string) {
	fmt.Println(C(BOLD+text, BMAGENTA))
}
