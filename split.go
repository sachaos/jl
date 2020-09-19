package main

import (
	"regexp"
	"strings"
)

var space = regexp.MustCompile("[[:blank:]]+")

func Split(s string) []string {
	s = strings.TrimSpace(s)
	s = space.ReplaceAllString(s, " ")
	return strings.Split(s, " ")
}
