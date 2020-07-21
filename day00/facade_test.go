package day00

import (
	"testing"
)

func TestFacade(t *testing.T) {
	expectedRes := "Simple text1 to show in example\nSimple text2 to show in example\nSimple text3 to show in example"

	example := InitActions()

	res := example.DoActions()

	if res != expectedRes {
		t.Errorf("Expect result to be : %s, but : %s.\n", expectedRes, res)
	}

}
