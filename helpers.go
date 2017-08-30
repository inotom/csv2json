package main

import (
	"strconv"
	"strings"
)

func string2Int(s string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(s), 10, 0)
}

func string2Float(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}
