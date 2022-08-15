package scenarios

import (
	"fmt"
	"path/filepath"

	"github.com/dv4mp1r3/ovpngen/common"
)

type ConfGenScenario interface {
	Scenario
}

type ConfGenScenarioImpl struct {
	ResultPath string
	Port       string
	Proto      string
	Addr       string
}

const (
	confTemplate string = `port %s
proto %s
dev tun
ca ca.crt
cert server.crt
key server.key
dh dh.pem
server %s 255.255.255.0
ifconfig-pool-persist /var/log/openvpn/ipp.txt
allow-pull-fqdn
keepalive 10 120
tls-auth ta.key 0 
cipher AES-256-CBC
persist-key
persist-tun
status /var/log/openvpn/openvpn-status.log
verb 3`
	ScenarioConfGen string = "confgen"
)

func (s *ConfGenScenarioImpl) Execute() {

	config := fmt.Sprintf(confTemplate, s.Port, s.Proto, s.Addr)
	if s.ResultPath == "" {
		fmt.Println(config)
	}

}

func (s *ConfGenScenarioImpl) Validate() bool {
	return s.Addr != "" && s.Port != "" && s.Proto != ""
}

func (s *ConfGenScenarioImpl) ShowUsage() {
	fmt.Println("Usage:")
	exe := common.GetExecFileName()

	optionsTemplate := `%s [options]

Options:
	-a %s
	-pr %s
	-p %s

Example: 
ovpngen -scn %s -a 127.0.0.1 -pr udp -p 443
`

	fmt.Printf(
		optionsTemplate,
		filepath.Base(exe),
		"IP address",
		"Protocol (tcp or udp)",
		"Port",
		ScenarioConfGen,
	)
}
