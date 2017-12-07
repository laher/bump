package main

import "testing"

var bumptests = []struct {
	V      string
	Params bumpParams
	out    string
	err    error
}{
	{"1", bumpParams{Part: 0, Inc: 1}, "2", nil},
	{"1.1", bumpParams{Part: 0, Inc: 1}, "1.2", nil},
	{"1.1.1", bumpParams{Part: 0, Inc: 1}, "1.1.2", nil},
	{"1.1.1", bumpParams{Part: 2, Inc: 1, LeftToRight: true}, "0.0.2", nil},
	{"1.1.1", bumpParams{Part: 1, Inc: 1}, "1.2.0", nil},
	{"1.1.1", bumpParams{Part: 2, Inc: 1}, "2.0.0", nil},
	{"app-v1.1.1", bumpParams{Prefix: "app-v", Part: 2, Inc: 1}, "app-v2.0.0", nil},
	{"1.1.1-xxx", bumpParams{Part: 0, Inc: 1}, "1.1.2", nil},
	{"1.1.xxx-1", bumpParams{Part: 0, Inc: 1}, "1.1.xxx-2", nil},
	{"1.1.xxx-", bumpParams{Part: 0, Inc: 1}, "", errNonNumeric},
	{"1", bumpParams{Part: 1, Inc: 1}, "", errInvalidPartNum},
	{"1", bumpParams{Part: -1, Inc: 1}, "", errInvalidPartNum},
	{"", bumpParams{}, "", errNoVersionSupplied},
}

func TestBump(t *testing.T) {
	for _, bt := range bumptests {
		t.Run(bt.V, func(t *testing.T) {
			t.Logf("Input: %+v", bt)
			v, err := toVersion(bt.V, &bt.Params)
			if err != nil {
				if err == bt.err {
					//done.
					return
				}
				t.Errorf("Fail: expected: %v, actual: %v", bt.err, err)
			}
			out, err := bump(v, bt.Params)
			if err != bt.err {
				t.Errorf("Fail: expected: %v, actual: %v", bt.err, err)
			}
			if out != bt.out {
				t.Errorf("Fail: expected: %s, actual: %s", bt.out, out)
			}
		})
	}
}
