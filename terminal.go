package main
import "fmt"
func SetTitle(
	text string,
) {
	fmt.Printf(
		"\033]0;%s\007",
		text,
	)
}
func ClearTitle() {
	fmt.Print(
		"\033]0;\007",
	)
}
func Bell() {
	fmt.Print(
		"\a",
	)
}
