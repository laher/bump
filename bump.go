package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var errInvalidPartNum = errors.New("version 'part' number invalid")
var errNonNumeric = errors.New("version contains a non-numeric part")
var errNoVersionSupplied = errors.New("empty version string")
var errNoPrefix = errors.New("Prefix not present")

type bumpParams struct {
	Part        int
	LeftToRight bool
	Delimiter   string
	Prefix      string
	Inc         int
	Sort        string
}

func toVersion(versionString string, params *bumpParams) (Version, error) {
	v := Version{}
	versionString = strings.TrimSpace(versionString)
	if versionString == "" {
		return v, errNoVersionSupplied
	}
	if params.Delimiter == "" {
		params.Delimiter = "."
	}
	r, err := regexp.Compile("^([^0-9]*)([\\d]+)(.*)")
	if err != nil {
		return v, err
	}

	if params.Prefix != "" {
		if !strings.HasPrefix(versionString, params.Prefix) {
			return v, errNoPrefix
		}
		versionString = versionString[len(params.Prefix):]
	}
	vparts := strings.Split(versionString, params.Delimiter)
	//fmt.Printf("parts: %v\n", vparts)
	for _, p := range vparts {
		subMatches := r.FindAllStringSubmatch(p, -1)
		//fmt.Printf("subMatches: %v\n", subMatches)
		if subMatches == nil {
			if params.Prefix == "" {
				return v, errNonNumeric
			}
			//prefix suggests we should stretch
			continue
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

func bump(version Version, params bumpParams) (string, error) {
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
	for i := range version.parts {
		if i == index {
			version.parts[i].val += params.Inc
			version.parts[i].suffix = ""
		} else if (params.LeftToRight && i < index) || (!params.LeftToRight && i > index) {
			version.parts[i].val = 0
			version.parts[i].suffix = ""
			version.parts[i].prefix = ""
		}
	}
	return version.ToString(params), nil
}
