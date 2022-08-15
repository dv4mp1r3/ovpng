package scenarios

import (
	"fmt"
	"path/filepath"

	"github.com/dv4mp1r3/ovpngen/common"
)

type Ovpngen interface {
	Scenario
}

type OvpngenImpl struct {
	args        []string
	RootDir     string
	Host        string
	Port        string
	CertPath    string
	KeyPath     string
	CaPath      string
	TlsAuthPath string
}

const (
	HelpStr string = "Show usage"

	ScenarioOvpnGenName string = "ovpngen"
)

func (s *OvpngenImpl) Execute() {
	ovpnTemplate := `client
	dev tun
	proto udp
	remote %s %s
	nobind
	persist-key
	persist-tun
	comp-lzo
	<ca>
	%s
	</ca>
	<cert>
	%s
	</cert>
	<key>
	%s
	</key>
	key-direction 1
	<tls-auth>
	%s
	</tls-auth>`
	if common.NeedToShowUsage() {
		s.ShowUsage()
		return
	}

	args := []string{s.RootDir, s.Host, s.Port, s.CertPath, s.KeyPath}
	if !common.ValidateArgs(args) {
		s.ShowUsage()
		return
	}

	fmt.Printf(
		ovpnTemplate,
		s.Host,
		s.Port,
		common.ReadFile(s.RootDir, s.CaPath),
		common.ReadFile(s.RootDir, s.CertPath),
		common.ReadFile(s.RootDir, s.KeyPath),
		common.ReadFile(s.RootDir, s.TlsAuthPath),
	)
}

func (s *OvpngenImpl) ShowUsage() {
	fmt.Println("Usage:")
	exe := common.GetExecFileName()

	optionsTemplate := `%s [options]

Options:
	-r %s
	-s %s
	-p %s
	-c %s
	-k %s
	-ca %s
	-ta %s

Example: 
ovpngen -scn %s -r /etc/openvpn -c easy-rsa/pki/issued/user.crt -k easy-rsa/pki/private/user.key -s 127.0.0.1 -p 1194
`
	fmt.Printf(
		optionsTemplate,
		filepath.Base(exe),
		common.RHelpStr,
		common.SHelpStr,
		common.PHelpStr,
		common.CertPathHelpStr,
		common.KeyPathHelpStr,
		common.CAPathHelpStr,
		common.TLSAuthPathHelpStr,
		ScenarioOvpnGenName,
	)
}

func (s *OvpngenImpl) Validate() bool {
	args := []string{s.RootDir, s.Host, s.Port, s.CertPath, s.KeyPath}
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}
	return true
}
