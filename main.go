package main

import (
	"flag"
	"fmt"

	"github.com/dv4mp1r3/ovpngen/common"
	"github.com/dv4mp1r3/ovpngen/scenarios"
)

const ()

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

	fmt.Println(fmt.Sprint("%s %s", certType, certName))

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
