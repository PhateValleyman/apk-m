package main

import (
	"archive/zip"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const AAPT_TIMEOUT = 5 * time.Second

type APKInfo struct {
	Name    string
	Package string
	Version string
	Code    string
}

func ScanAPK(path string) {
	realAPK := path
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".apk" {
		tmp, ok := ExtractBaseAPK(path)
		if !ok {
			fmt.Println(C("Cannot extract "+path, RED))
			return
		}
		realAPK = tmp
	}
	typ := "APK"
	switch ext {
	case ".apks":
		typ = "APKS"
	case ".xapk":
		typ = "XAPK"
	case ".apkm":
		typ = "APKM"
	}
	info, ok := ReadAPKInfo(realAPK)
	if !ok {
		fmt.Println(C("[TIMEOUT/ERROR] ", RED), path)
		SaveRecord(path, CurrentConfig.StartPath, typ, "", "", "", "", SHA256(path), 0, "", "UNREADABLE", 0)
		return
	}
	hash := SHA256(path)
	size := int64(0)
	stat, err := os.Stat(path)
	if err == nil {
		size = stat.Size()
	}
	signature, status, isMod := CheckSignature(realAPK)
	SaveRecord(path, CurrentConfig.StartPath, typ, info.Name, info.Package, info.Version, info.Code, hash, size, signature, status, isMod)
	fmt.Println(C("[FOUND] ", GREEN), info.Package, info.Version)
}
func ExtractBaseAPK(path string) (string, bool) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return "", false
	}
	defer reader.Close()
	tmp := filepath.Join(os.TempDir(), "apk-m-base.apk")
	for _, file := range reader.File {
		if file.Name == "base.apk" {
			in, err := file.Open()
			if err != nil {
				return "", false
			}
			out, err := os.Create(tmp)
			if err != nil {
				return "", false
			}
			io.Copy(out, in)
			out.Close()
			in.Close()
			return tmp, true
		}
	}
	return "", false
}
func ReadAPKInfo(path string) (APKInfo, bool) {
	result := APKInfo{}
	out, err := runAapt("aapt", path)
	if err != nil {
		out, err = runAapt("aapt2", path)
		if err != nil {
			return result, false
		}
	}
	text := string(out)
	result.Name = findAPK(`label='([^']+)'`, text)
	result.Package = findAPK(`name='([^']+)'`, text)
	result.Version = findAPK(`versionName='([^']+)'`, text)
	result.Code = findAPK(`versionCode='([^']+)'`, text)
	if result.Package == "" {
		return result, false
	}
	return result, true
}
func runAapt(bin string, path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), AAPT_TIMEOUT)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "dump", "badging", path)
	out, err := cmd.Output()
	if ctx.Err() == context.DeadlineExceeded {
		return out, context.DeadlineExceeded
	}
	return out, err
}
func isTimeout(err error) bool {
	return err == context.DeadlineExceeded
}
func findAPK(pattern string, text string) string {
	r := regexp.MustCompile(pattern)
	m := r.FindStringSubmatch(text)
	if len(m) > 1 {
		return m[1]
	}
	return ""
}
func SHA256(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	hash := sha256.New()
	io.Copy(hash, file)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
