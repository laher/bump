package main

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	name   string
	in     RSorted
	sorted RSorted
}{
	{
		name:   "single",
		in:     RSorted{Version{parts: []part{{val: 1}, {val: 1}}}},
		sorted: RSorted{Version{parts: []part{{val: 1}, {val: 1}}}},
	},
	{
		name:   "2 RSorted",
		in:     RSorted{Version{parts: []part{{val: 1}}}, Version{parts: []part{{val: 2}}}},
		sorted: RSorted{Version{parts: []part{{val: 2}}}, Version{parts: []part{{val: 1}}}},
	},
}

func TestSort(t *testing.T) {

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			t.Logf("unsorted %v", tst.in)
			sort.Sort(tst.in)
			if !reflect.DeepEqual(tst.in, tst.sorted) {
				t.Errorf("%v is not equal to expected %v", tst.in, tst.sorted)
			}
		})
	}

}
