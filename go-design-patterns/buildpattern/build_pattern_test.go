package buildpackage

import "testing"

func TestBuildPattern(t *testing.T) {
	creator := Creator{
		computer: HuaWeiComputer{},
	}
	creator.Construct()
	t.Logf("%v",creator)
}
