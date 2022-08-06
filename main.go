package main

import (
	"flag"
	"fmt"

	"github.com/dv4mp1r3/ovpngen/common"
	"github.com/dv4mp1r3/ovpngen/scenarios"
)

func main() {

	scenario := flag.String("scn", "", "Scenario")
	flag.Parse()
	if len(*scenario) == 0 {
		showUsage()
		return
	}
	switch *scenario {
	case scenarios.ScenarioOvpnGenName:
		var s = new(scenarios.OvpngenImpl)
		s.Execute()
	case scenarios.ScenarioEasyRsaName:
		var s = new(scenarios.EasyRsaScenarioImpl)
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
