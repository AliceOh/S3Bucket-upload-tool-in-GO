package main

import (
	"flag"
	"fmt"
)

var versionFlags = flag.NewFlagSet(appName, flag.ContinueOnError)

func getVersion() string {
	return fmt.Sprintf("%s [%s] %s", appName, appDescription, appVersion)
}

func version(argv []string) int {
	// Parse arguments.
	if err := versionFlags.Parse(argv); err != nil {
		fmt.Println(err.Error())
		fmt.Println(helpRoot)
		return 1
	}
	// Print version
	fmt.Println(getVersion())
	return 0
}
