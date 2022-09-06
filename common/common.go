package common

import (
	"os"
)

const (
	CreateClientKey    string = "client-key"
	CreateServerKey    string = "server-key"
	CaPwd              string = "CA password"
	ServerPwd          string = "Server's certificate password"
	SHelpStr           string = "server's host"
	PHelpStr           string = "server's port"
	CertPathHelpStr    string = "user's .crt file (inside root)"
	KeyPathHelpStr     string = "user's .key file (inside root)"
	CAPathHelpStr      string = "CA path (inside root, ca.crt by default)"
	TLSAuthPathHelpStr string = "TLS auth path (inside root, ta.key by default)"
	CertTypeHelpStr    string = "Certificate type (%s or %s)"
	CertNameHelpStr    string = "Certificate name"
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

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
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
