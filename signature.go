package main

import (
	"os/exec"
	"strings"
)

func CheckSignature(path string) (string, string, int) {
	cmd := exec.Command(
		"apksigner",
		"verify",
		"--print-certs",
		path,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", "UNSIGNED", 1
	}
	text := string(out)
	signature := ""
	for _, line := range strings.Split(
		text,
		"\n",
	) {
		if strings.Contains(
			line,
			"SHA-256",
		) {
			signature = line
		}
	}
	status := "ORIGINAL"
	isMod := 0
	lower := strings.ToLower(text)
	if strings.Contains(
		lower,
		"android debug",
	) {
		status = "DEBUG"
		isMod = 1
	}
	return signature, status, isMod
}
