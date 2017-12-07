package bump

import (
	"fmt"
	"strings"
)

// Version is a simple representation of a version without params
type Version struct {
	parts []part
}

type part struct {
	val    int
	prefix string
	suffix string
}

func (p part) String() string {
	return fmt.Sprintf("%s%d%s", p.prefix, p.val, p.suffix)
}

// ToString renders a version based on its params
func (v Version) ToString(params BumpParams) string {
	vs := make([]string, len(v.parts))
	for i, p := range v.parts {
		vs[i] = p.String()
	}
	vSt := strings.Join(vs, params.Delimiter)
	return params.Prefix + vSt
}

// Sorted is a slice of versions, for sorting in ascending order (Lowest first)
type Sorted []Version

func (v Sorted) Len() int      { return len(v) }
func (v Sorted) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v Sorted) Less(i, j int) bool {
	if len(v[i].parts) > len(v[j].parts) {
		return false
	}
	if len(v[i].parts) < len(v[j].parts) {
		return true
	}
	for k, part := range v[i].parts {
		if part.val != v[j].parts[k].val {
			return part.val < v[j].parts[k].val
		}
	}
	return true
}

// RSorted is a slice of versions, for sorting in reverse-order (Highest first)
type RSorted []Version

func (v RSorted) Len() int      { return len(v) }
func (v RSorted) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v RSorted) Less(i, j int) bool {
	if len(v[i].parts) < len(v[j].parts) {
		return false
	}
	if len(v[i].parts) > len(v[j].parts) {
		return true
	}
	for k, part := range v[i].parts {
		if part.val != v[j].parts[k].val {
			return part.val > v[j].parts[k].val
		}
	}
	return true
}
