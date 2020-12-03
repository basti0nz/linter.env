package main

import (
    "github.com/pborman/getopt/v2"
    "fmt"
    "github.com/versus/linter.env/cmd"
)

func main() {
    showChecksPtr := getopt.BoolLong("show-checks", 'l', "view all available checks")
    helpFlag := getopt.BoolLong("help", 'h', "Displays help message")
    excludeFlag := getopt.StringLong("exclude", 'e', "", "for exclude a file from check")
    recursiveFlag := getopt.BoolLong("recursive", 'r', "recursive search inside directories ")
    skipRuleFlag := getopt.StringLong("skip", 's', "","skip some checks")
    quietFlag := getopt.BoolLong("quiet", 'q', "see only warnings without additional information")
    noBackupFlag := getopt.BoolLong("no-backup", 'b',   "disable the backup function")
    fixFlag := getopt.BoolLong("fix", 'f',   "automatically fix warnings in the files")

    getopt.Parse()
    args := getopt.Args()
    if len(args) > 1 || *helpFlag{
		cmd.UsageMessage()
	}


    if *showChecksPtr {
        cmd.ShowRules()
	}

    fmt.Println(*excludeFlag, "tail:", args, *recursiveFlag, *skipRuleFlag, *quietFlag, *noBackupFlag, *fixFlag)
}