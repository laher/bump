package bump

import "testing"

var bumptests = []struct {
	Params BumpParams
	out    string
	err    error
}{
	{BumpParams{"1", 0, false}, "2", nil},
	{BumpParams{"1.1", 0, false}, "1.2", nil},
	{BumpParams{"1.1.1", 0, false}, "1.1.2", nil},
	{BumpParams{"1.1.1", 2, true}, "1.1.2", nil},
	{BumpParams{"1.1.1", 1, false}, "1.2.0", nil},
	{BumpParams{"1.1.1", 2, false}, "2.0.0", nil},
	{BumpParams{"1.1.1-xxx", 0, false}, "1.1.2", nil},
	{BumpParams{"1.1.xxx-1", 0, false}, "1.1.xxx-2", nil},
	{BumpParams{"1.1.xxx-", 0, false}, "", errNonNumeric},
	{BumpParams{"1", 1, false}, "", errInvalidPartNum},
	{BumpParams{"1", -1, false}, "", errInvalidPartNum},
}

func TestBump(t *testing.T) {
	for _, bt := range bumptests {
		t.Logf("table driven test: %+v", bt)
		out, err := Bump(bt.Params)
		if err != bt.err {
			t.Errorf("Fail: expected: %v, actual: %v", bt.err, err)
		}
		if out != bt.out {
			t.Errorf("Fail: expected: %s, actual: %s", bt.out, out)
		}
	}
}
