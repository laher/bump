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
}

func Bump(params BumpParams) (string, error) {
	v := strings.TrimSpace(params.V)
	if v == "" {
		return "", errNoVersionSupplied
	}
	if params.Delimiter == "" {
		params.Delimiter = "."
	}
	vparts := strings.Split(v, params.Delimiter)
	if params.Part < 0 {
		return "", errInvalidPartNum
	}
	max := len(vparts) - 1
	if params.Part > max {
		return "", errInvalidPartNum
	}
	index := params.Part
	if !params.LeftToRight {
		index = max - params.Part
	}
	thisPart := vparts[index]
	r, err := regexp.Compile("^([^0-9]*)([\\d+])(.*)")
	if err != nil {
		return "", err
	}
	thisPartPrefix := ""
	thisPartInt, err := strconv.Atoi(thisPart)
	if err != nil {
		subMatches := r.FindAllStringSubmatch(thisPart, -1)
		if subMatches == nil {
			return "", errNonNumeric
		}
		sm0 := subMatches[0]
		thisPartNumeric := sm0[2]
		if len(thisPartNumeric) < 1 {
			return "", errNonNumeric
		}
		thisPartInt, err = strconv.Atoi(thisPartNumeric)
		if err != nil {
			return "", err
		}
		thisPartPrefix = sm0[1]
	}
	thisPartInt += 1
	vparts[index] = thisPartPrefix + strconv.Itoa(thisPartInt)
	for i, _ := range vparts[index+1:] {
		vparts[i+index+1] = "0"
	}
	vNew := strings.Join(vparts, params.Delimiter)
	return vNew, nil
}
