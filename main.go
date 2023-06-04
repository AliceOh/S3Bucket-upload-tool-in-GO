package main

import (
	"flag"
	"fmt"
	"os"
)

// app version information
var appName = os.Args[0]
var appDescription = "Iress Content Team S3 Backup Tool"
var appVersion = "1.0.0"

// Runtime flags.
var flags = flag.NewFlagSet(appName, flag.ContinueOnError)
var flagHelp = flags.Bool("help", false, "")

// Help string.
var helpCommands = `commands:

  uploads3         upload an environment configuration file
  version           display version`

var helpGlobalOptions = `global options:

  --help          display help`

var helpRoot = fmt.Sprintf(
	"usage: %s [global options...] <command> [command options...]\n\n%s\n\n%s",
	appName,
	helpCommands,
	helpGlobalOptions,
)

func main() {
	os.Exit(func() int {
		// Parse arguments.
		if err := flags.Parse(os.Args[1:]); err != nil {
			fmt.Println(err.Error())
			fmt.Println(helpRoot)
			return 1
		}

		// Determine the action based on the command.
		switch flags.Arg(0) {
		case "version":
			return version(flags.Args()[1:])
		case "uploads3":
			return uploads3(flags.Args()[1:])
		default:
			if *flagHelp {
				fmt.Println(helpRoot)
				return 0
			} else {
				fmt.Println("invalid command")
				fmt.Println(helpRoot)
				return 1
			}
		}
	}())
}
