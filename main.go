package main

import (
	"flag"
	"fmt"

	"github.com/dv4mp1r3/ovpngen/common"
	"github.com/dv4mp1r3/ovpngen/scenarios"
)

func main() {

	//todo: change args
	scenario := flag.String("scn", "", "Scenario")
	help := flag.Bool("h", false, "Show usage")
	certType := flag.String(
		"ct",
		"",
		fmt.Sprint(
			"Certificate type (%s, %s)",
			common.CreateServerKey,
			common.CreateClientKey,
		),
	)
	certName := flag.String("n", "", "Certificate name")
	cP := flag.String("cp", "", common.CaPwd)
	sP := flag.String("sp", "", common.ServerPwd)
	//===
	root := flag.String("r", "", common.RHelpStr)
	host := flag.String("s", "", common.SHelpStr)
	port := flag.String("p", "", common.PHelpStr)
	certPath := flag.String("c", "", common.CertPathHelpStr)
	keyPath := flag.String("k", "", common.KeyPathHelpStr)
	caPath := flag.String("ca", "ca.crt", common.CAPathHelpStr)
	tlsAuthPath := flag.String("ta", "ta.key", common.TLSAuthPathHelpStr)
	//==
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
		if *help {
			s.ShowUsage()
			return
		}
		s.RootDir = *root
		s.Host = *host
		s.Port = *port
		s.CertPath = *certPath
		s.KeyPath = *keyPath
		s.CaPath = *caPath
		s.TlsAuthPath = *tlsAuthPath
		s.Execute()
	case scenarios.ScenarioEasyRsaName:
		var s = new(scenarios.EasyRsaScenarioImpl)
		s.CertName = *certName
		s.CertType = *certType
		s.CaPwd = cP
		s.CertPwd = sP
		if s.CertName == "" {
			fmt.Println("Certificate name can not be empty")
			return
		}
		s.Execute()
	case scenarios.ScenarioConfGen:
		var s = new(scenarios.ConfGenScenarioImpl)
		s.Addr = *addr
		s.Port = *port
		s.Proto = *proto
		if !s.Validate() {
			fmt.Printf("Some of parameters (proto, port, addr) is empty.")
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
