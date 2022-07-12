package core

import "testing"

func TestScenario(t *testing.T) {
	scenarioName := "test"
	s := Scenario(scenarioName)

	if s.name != scenarioName {
		t.FailNow()
	}

}
