package filehandler

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//如果是如下路径，则跳过
var skipDir = []string{
	"/.git",
	"/.vscode",
	"/vendor",
	"/_output",
	"/generated",
}

//需要跳过的文件
var skipFile = []string{
	"_test.cc",
	"_test.go",
	"Test.java",
}

// 如果含有如下字符串，则跳过该行
var skipString = []string{
	"```",
	`{%`,
	`%}`,
	`{{`,
}

//需要处理的文件类型
var fileType string
//两个相同行的最大行距
var lineGap int

//入口函数
func ScanRun(dir string) {
	fmt.Println(fileType)
	err := filepath.Walk(dir, walkHandler)
	if err != nil {
		panic(err)
	}
}

//walk的处理函数
func walkHandler(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("error occurred, stop.")
		return err
	}

	if info.IsDir() {
		for _, s := range skipDir {
			if strings.Contains(path, s) {
				//fmt.Println("skip path")
				//fmt.Println(path)
				return filepath.SkipDir
			}
		}
		//fmt.Println("is dir, no action")
		//fmt.Println(path)
		return nil
	}
	if isTarget(path) {
		if ferr := FileScanner(path); ferr != nil {
			return ferr
		}
		//fmt.Println("action file")
	}
	return nil
}

//检查文件是否需要扫描
func isTarget(path string) bool {
	for _, s := range skipFile {
		if strings.Contains(path, s) {
			return false
		}
	}
	ext := filepath.Ext(path)
	ft := strings.ToLower(fileType)
	switch ft {
	case "md":
		if strings.EqualFold(ext, ".md") {
			return true
		}
		return false
	case "go":
		if strings.EqualFold(ext, ".go") {
			return true
		}
		return false
	case "py":
		if strings.EqualFold(ext, ".py") {
			return true
		}
		return false
	case "java":
		if strings.EqualFold(ext, ".java") {
			return true
		}
		return false
	case "yaml", "yml":
		if strings.EqualFold(ext, ".yaml") || strings.EqualFold(ext, ".yml") {
			return true
		}
		return false
	case "all":
		return true
	default:
		return false
	}
}

//扫描文件
func FileScanner(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer file.Close()

	cache := make(map[string]int)
	lineNum := 1
	var line string
	var firstPrint = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if isSkipLine(line) {
			lineNum++
			continue
		}
		if _, isIn := cache[line]; isIn {
			if lineNum - cache[line] <= lineGap {
				if firstPrint {
					fmt.Println()
					fmt.Println(path)
					firstPrint = false
				}
				fmt.Printf("line number: %d %d\n", cache[line], lineNum)
				fmt.Println(line)
			}
		}
		cache[line] = lineNum
		lineNum++
	}
	//fmt.Println(path)
	return nil
}

//是否需要跳过该行
func isSkipLine(line string) bool {
	if line == "" {
		return true
	}
	if len(line) <= 7 {
		return true
	}
	for _, s := range skipString {
		if strings.Contains(line, s) {
			return true
		}
	}
	//不包含字母的行不检测
	if match, _ := regexp.MatchString(`[a-zA-Z]+.* .*[a-zA-Z]+`, line); !match {
		return true
	}

	return false
}

func init() {
	flag.StringVarP(&fileType, "type", "t", "all", "the file type to examine")
	flag.IntVarP(&lineGap, "gap", "g", 1, "the max line gap to examine")
}
