package scenarios

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/dv4mp1r3/ovpngen/common"
)

type Ovpngen interface {
	Scenario
}

type OvpngenImpl struct {
	args []string
}

const (
	RHelpStr           string = "openvpn config root path"
	SHelpStr           string = "server's host"
	PHelpStr           string = "server's port"
	CertPathHelpStr    string = "user's .crt file (inside root)"
	KeyPathHelpStr     string = "user's .key file (inside root)"
	CAPathHelpStr      string = "CA path (inside root, ca.crt by default)"
	TLSAuthPathHelpStr string = "TLS auth path (inside root, ta.key by default)"
	HelpStr            string = "Show usage"

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

	root := flag.String("r", "", RHelpStr)
	host := flag.String("s", "", SHelpStr)
	port := flag.String("p", "", PHelpStr)
	certPath := flag.String("c", "", CertPathHelpStr)
	keyPath := flag.String("k", "", KeyPathHelpStr)
	caPath := flag.String("ca", "ca.crt", CAPathHelpStr)
	tlsAuthPath := flag.String("ta", "ta.key", TLSAuthPathHelpStr)

	flag.Parse()
	args := []string{*root, *host, *port, *certPath, *keyPath}
	if !common.ValidateArgs(args) {
		s.ShowUsage()
		return
	}

	fmt.Printf(
		ovpnTemplate,
		*host,
		*port,
		common.ReadFile(*root, *caPath),
		common.ReadFile(*root, *certPath),
		common.ReadFile(*root, *keyPath),
		common.ReadFile(*root, *tlsAuthPath),
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
		RHelpStr,
		SHelpStr,
		PHelpStr,
		CertPathHelpStr,
		KeyPathHelpStr,
		CAPathHelpStr,
		TLSAuthPathHelpStr,
		ScenarioOvpnGenName,
	)
}
