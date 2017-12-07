package bump

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var errInvalidPartNum = errors.New("version 'part' number invalid")
var errNonNumeric = errors.New("version contains a non-numeric component")
var errNoVersionSupplied = errors.New("empty version string")

type BumpParams struct {
	V           string
	Part        int
	LeftToRight bool
	Delimiter   string
	Prefix      string
	Amount      int
	Sort        string
}

func ToVersion(versionString string, params BumpParams) (Version, error) {

	v := Version{}
	r, err := regexp.Compile("^([^0-9]*)([\\d]+)")
	if err != nil {
		return v, err
	}
	if strings.HasPrefix(versionString, params.Prefix) {
		versionString = versionString[len(params.Prefix):]
	}
	vparts := strings.Split(versionString, params.Delimiter)
	//fmt.Printf("parts: %v\n", vparts)
	for _, p := range vparts {
		subMatches := r.FindAllStringSubmatch(p, -1)
		//fmt.Printf("subMatches: %v\n", subMatches)
		if subMatches == nil {
			return v, errNonNumeric
		}
		sm0 := subMatches[0]
		thisPartPrefix := sm0[1]
		thisPartNumeric := sm0[2]
		thisPartSuffix := ""
		if len(sm0) > 3 {
			thisPartSuffix = sm0[3]
		}
		if len(thisPartNumeric) < 1 {
			return v, errNonNumeric
		}
		thisPartInt, err := strconv.Atoi(thisPartNumeric)
		if err != nil {
			return v, err
		}

		v.parts = append(v.parts, part{prefix: thisPartPrefix, val: thisPartInt, suffix: thisPartSuffix})
	}

	return v, nil
}

func Bump(version Version, params BumpParams) (string, error) {

	//vparts := strings.Split(v, params.Delimiter)
	if params.Part < 0 {
		return "", errInvalidPartNum
	}
	max := len(version.parts) - 1
	if params.Part > max {
		return "", errInvalidPartNum
	}
	index := params.Part
	if !params.LeftToRight {
		index = max - params.Part
	}
	thisPart := version.parts[index]
	thisPart.val += params.Amount
	return version.ToString(params), nil
}
