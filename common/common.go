package common

import (
	"os"
)

const (
	RHelpStr string = "openvpn config root path"
)

func NeedToShowUsage() bool {
	for idx, arg := range os.Args {
		if (arg == "-h" || arg == "--help") && idx%2 == 1 {
			return true
		}
	}
	return false
}

func ValidateArgs(args []string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return false
		}
	}
	return true
}

func ReadFile(root string, path string) string {
	data, err := os.ReadFile(root + string(os.PathSeparator) + path)
	checkFileErr(err)
	return string(data)
}

func checkFileErr(e error) {
	if e != nil {
		panic(e)
	}
}

func GetExecFileName() string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return exe
}
