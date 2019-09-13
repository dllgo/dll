package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type FormatData struct {
	PkgName string
	Name    string
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

func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		return filepath.Join(home, "go")
	}
	return ""
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" && strings.Compare(runtime.Version(), "go1.8") >= 0 {
		gopath = defaultGOPATH()
	}
	return filepath.SplitList(gopath)
}

// IsInGOPATH checks whether the path is inside of any GOPATH or not
func IsInGOPATH(thePath string) bool {
	for _, gopath := range GetGOPATHs() {
		if strings.Contains(thePath, filepath.Join(gopath, "src")) {
			return true
		}
	}
	return false
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// WriteToFile creates a file and writes content to it
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}
func CheckEnv(appname string) (apppath, packpath string, err error) {
	gps := GetGOPATHs()
	if len(gps) == 0 {
		fmt.Println("GOPATH environment variable is not set or empty")
	}
	currpath, _ := os.Getwd()
	currpath = filepath.Join(currpath, appname)
	for _, gpath := range gps {
		gsrcpath := filepath.Join(gpath, "src")
		if strings.HasPrefix(strings.ToLower(currpath), strings.ToLower(gsrcpath)) {
			packpath = strings.Replace(currpath[len(gsrcpath)+1:], string(filepath.Separator), "/", -1)
			return currpath, packpath, nil
		}
	}

	// In case of multiple paths in the GOPATH, by default
	// we use the first path
	gopath := gps[0]

	fmt.Println("You current workdir is not inside $GOPATH/src.")

	gosrcpath := filepath.Join(gopath, "src")
	apppath = filepath.Join(gosrcpath, appname)

	if _, e := os.Stat(apppath); !os.IsNotExist(e) {
		err = fmt.Errorf("Cannot create application without removing '%s' first", apppath)
		fmt.Println("Path '%s' already exists", apppath)
		return
	}
	packpath = strings.Join(strings.Split(apppath[len(gosrcpath)+1:], string(filepath.Separator)), "/")
	return
}
