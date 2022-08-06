package scenarios

import "fmt"

type EasyRsaScenario interface {
	Scenario
}

type EasyRsaScenarioImpl struct {
	args []string
}

const (
	ScenarioEasyRsaName string = "easyrsa"
)

func (s *EasyRsaScenarioImpl) Execute() {
	fmt.Println("Not implemented yet :(")
}
