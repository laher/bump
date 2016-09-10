package bump

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var errInvalidPartNum = errors.New("version 'part' number invalid")
var errNonNumeric = errors.New("version contains a non-numeric component")

type BumpParams struct {
	V           string
	Part        int
	LeftToRight bool
}

func Bump(params BumpParams) (string, error) {
	vparts := strings.Split(params.V, ".")
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
		/*
			_, err := strconv.Atoi(p)
			if err != nil {
				break
			} else {
				//reset smaller parts to 0
			}
		*/
		vparts[i+index+1] = "0"
	}
	vNew := strings.Join(vparts, ".")
	return vNew, nil
}
