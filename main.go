package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	NsFactor          int64 = 10000000
	SecsBetweenEpochs int64 = 11644473600
)

func boolify(val string) bool {
	switch val {
	case "1":
		return true
	case "true":
		return true
	case "yes":
		return true
	case "y":
		return true
	default:
		return false
	}
}

// Convert filetime to time.Time
func FiletimeToDate(filetime int64) time.Time {
	// convert filetime to seconds, then subtract the seconds from 1601
	// to get epoch seconds
	return time.Unix(filetime/NsFactor-SecsBetweenEpochs, 0)
}

func main() {
	if boolify(os.Getenv("DEBUG")) {
		PrintBuildInfo()
	}
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <filetime>\n", os.Args[0])
		fmt.Printf("Set env var DEBUG to a truthy value to see build information\n")
		os.Exit(1)
	}
	filetime, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("error converting converting to integer: %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("%v\n", FiletimeToDate(filetime))
}
