package scenarios

type Scenario interface {
	Execute()
	ShowUsage()
}
