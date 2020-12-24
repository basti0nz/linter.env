package cmd

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"os"
)

func ShowRules()  {
	rules := []string {"DuplicatedKey",
	"EndingBlankLine",
	"ExtraBlankLine",
	"IncorrectDelimiter",
	"KeyWithoutValue",
	"LeadingCharacter",
	"LowercaseKey",
	"QuoteCharacter",
	"SpaceCharacter",
	"SpecialCharacter",
	"TrailingWhitespace",
	"UnorderedKey" }

	fmt.Println( rules)
}

func UsageMessage()  {
	getopt.PrintUsage(os.Stderr)
	os.Exit(0)
}
