package main

import (
	"flag"
	"fmt"

	"github.com/dv4mp1r3/ovpngen/common"
	"github.com/dv4mp1r3/ovpngen/scenarios"
)

func main() {

	//common args
	scenario := flag.String("scn", "", "Scenario")
	help := flag.Bool("h", false, "Show usage")
	port := flag.String("p", "", common.PHelpStr)

	//easyrsa args
	certType := flag.String(
		"ct",
		"",
		fmt.Sprintf(
			"Certificate type (%s, %s)",
			common.CreateServerKey,
			common.CreateClientKey,
		),
	)
	certName := flag.String("n", "", "Certificate name")
	cP := flag.String("cp", "", common.CaPwd)
	sP := flag.String("sp", "", common.ServerPwd)

	//ovpngen args
	host := flag.String("s", "", common.SHelpStr)
	certPath := flag.String("c", "", common.CertPathHelpStr)
	keyPath := flag.String("k", "", common.KeyPathHelpStr)
	caPath := flag.String("ca", "ca.crt", common.CAPathHelpStr)
	tlsAuthPath := flag.String("ta", "ta.key", common.TLSAuthPathHelpStr)

	//confgen args
	addr := flag.String("a", "", "ip addr")
	proto := flag.String("pr", "", "proto")

	flag.Parse()

	if len(*scenario) == 0 {
		showUsage()
		return
	}
	switch *scenario {
	case scenarios.ScenarioOvpnGenName:
		var s = new(scenarios.OvpngenImpl)
		s.Host = *host
		s.Port = *port
		s.CertPath = *certPath
		s.KeyPath = *keyPath
		s.CaPath = *caPath
		s.TlsAuthPath = *tlsAuthPath
		if *help || !s.Validate() {
			s.ShowUsage()
			return
		}
		s.Execute()
	case scenarios.ScenarioEasyRsaName:
		var s = new(scenarios.EasyRsaScenarioImpl)
		s.CertName = *certName
		s.CertType = *certType
		s.CaPwd = cP
		s.CertPwd = sP
		if *help || !s.Validate() {
			s.ShowUsage()
			return
		}
		s.Execute()
	case scenarios.ScenarioConfGen:
		var s = new(scenarios.ConfGenScenarioImpl)
		s.Addr = *addr
		s.Port = *port
		s.Proto = *proto
		if *help || !s.Validate() {
			s.ShowUsage()
			return
		}
		s.Execute()
	}

}

func showUsage() {
	template := `Usage: %s -scn [scenario] <scenario params>

For more details see %s -scn [scenario] -h
`
	exe := common.GetExecFileName()
	fmt.Printf(template, exe, exe)
}
