package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NsFactor          int64 = 10000000
	SecsBetweenEpochs int64 = 11644473600
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
)

// GetBuildInfo returns a list of strings for either printing or logging
func GetBuildInfo() []string {
	longLine := strings.Replace(fmt.Sprintf("+%61s+", ""), " ", "-", 61)
	return []string{
		longLine,
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

// Convert filetime to time.Time
func FiletimeToDate(filetime int64) time.Time {
	// convert filetime to seconds, then subtract the seconds from 1601
	// to get epoch seconds
	return time.Unix(filetime/NsFactor-SecsBetweenEpochs, 0)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <filetime>\n", os.Args[0])
		os.Exit(1)
	}
	filetime, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("error converting converting to integer: %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("%v\n", FiletimeToDate(filetime))
}
