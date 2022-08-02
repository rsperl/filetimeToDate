package main

import (
	"fmt"
	"strings"
)

// Declare public variables to be set at build-time
var (
	BuildTime        string
	BuildCommit      string
	BuildShortCommit string
	BuildBranch      string
	BuildHostname    string
	BuildUsername    string
	Version          string
	Repository       string
)

// GetBuildInfo returns a list of strings for either printing or logging
func GetBuildInfo() []string {
	longLine := strings.Replace(fmt.Sprintf("+%61s+", ""), " ", "-", 61)
	return []string{
		longLine,
		fmt.Sprintf("| Repository:       %-41s |", Repository),
		fmt.Sprintf("| Version:          %-41s |", Version),
		fmt.Sprintf("| BuildTime:        %-41s |", BuildTime),
		fmt.Sprintf("| BuildCommit:      %-41s |", BuildCommit),
		fmt.Sprintf("| BuildShortCommit: %-41s |", BuildShortCommit),
		fmt.Sprintf("| BuildBranch:      %-41s |", BuildBranch),
		fmt.Sprintf("| BuildHostname:    %-41s |", BuildHostname),
		fmt.Sprintf("| BuildUsername:    %-41s |", BuildUsername),
		fmt.Sprintf(longLine),
	}
}

func PrintBuildInfo() {
	for _, line := range GetBuildInfo() {
		fmt.Println(line)
	}
}