package main
import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
func SetTitle(text string) {
	fmt.Printf("\033]0;%s\007", text)
}
func ClearTitle() {
	fmt.Print("\033]0;\007")
}
func Bell() {
	fmt.Print("\a")
}
func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80
	}
	parts := strings.Fields(string(out))
	if len(parts) < 2 {
		return 80
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 80
	}
	return width
}
