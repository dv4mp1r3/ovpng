package main

import (
	"flag"
	"fmt"

	"github.com/dv4mp1r3/ovpngen/common"
	"github.com/dv4mp1r3/ovpngen/scenarios"
)

func main() {

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
	}

}

func showUsage() {
	template := `Usage: %s -scn [scenario] <scenario params>

For more details see %s -scn [scenario] -h
`
	exe := common.GetExecFileName()
	fmt.Printf(template, exe, exe)
}
