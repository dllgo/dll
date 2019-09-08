package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type InputData struct {
	PkgName   string
	Name      string
	KebabName string // FuckShit -> fuck-shit
	// Fields    []GenerateField
}

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

// Converts a string to CamelCase
func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

// Converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}

/*
获取程序名称
*/
func StructName(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return ToCamel(args[0])
}

/*
获取程序包名
*/
func PkgName(args []string) string {
	if len(args) == 0 {
		return ""
	}
	if len(args) > 1 {
		return ToLowerCamel(args[1])
	}
	return ToLowerCamel(args[0])
}

/*
获取程序路径
*/
func PathArg(args []string) string {
	if len(args) > 2 {
		return args[2]
	}
	return GetCurrentDirectory()
}

/*
获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func GetFilePath(baseDir, sub string) (filepath string, err error) {
	filepath = path.Join(baseDir, sub)
	base := path.Dir(filepath)
	err = os.MkdirAll(base, os.ModePerm)
	return
}

// 写文件
func WriteFile(filepath string, content string) error {
	exists, err := PathExists(filepath)
	if err != nil {
		return err
	}
	if exists {
		fmt.Println("文件已经存在...", filepath)
		filepath = filepath + ".temp"
	}
	return WriteString(filepath, content, true)
}

func WriteString(path string, content string, append bool) error {
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag = flag | os.O_APPEND
	}
	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}
