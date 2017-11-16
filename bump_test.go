package bump

import "testing"

var bumptests = []struct {
	Params BumpParams
	out    string
	err    error
}{
	{BumpParams{V: "1", Part: 0}, "2", nil},
	{BumpParams{V: "1.1", Part: 0}, "1.2", nil},
	{BumpParams{V: "1.1.1", Part: 0}, "1.1.2", nil},
	{BumpParams{V: "1.1.1", Part: 2, LeftToRight: true}, "1.1.2", nil},
	{BumpParams{V: "1.1.1", Part: 1}, "1.2.0", nil},
	{BumpParams{V: "1.1.1", Part: 2}, "2.0.0", nil},
	{BumpParams{Prefix: "app-v", V: "app-v1.1.1", Part: 2}, "app-v2.0.0", nil},
	{BumpParams{V: "1.1.1-xxx", Part: 0}, "1.1.2", nil},
	{BumpParams{V: "1.1.xxx-1", Part: 0}, "1.1.xxx-2", nil},
	{BumpParams{V: "1.1.xxx-", Part: 0}, "", errNonNumeric},
	{BumpParams{V: "1", Part: 1}, "", errInvalidPartNum},
	{BumpParams{V: "1", Part: -1}, "", errInvalidPartNum},
	{BumpParams{V: ""}, "", errNoVersionSupplied},
}

func TestBump(t *testing.T) {
	for _, bt := range bumptests {
		t.Logf("Input: %+v", bt.Params)
		out, err := Bump(bt.Params)
		if err != bt.err {
			t.Errorf("Fail: expected: %v, actual: %v", bt.err, err)
		}
		if out != bt.out {
			t.Errorf("Fail: expected: %s, actual: %s", bt.out, out)
		}
	}
}
