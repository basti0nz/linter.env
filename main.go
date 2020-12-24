package main

import (
    "github.com/pborman/getopt/v2"
    "fmt"
    "github.com/versus/linter.env/cmd"
    "math/rand"
    "sync"
    "time"
    "github.com/BTBurke/clt"
)

func main() {
    showChecksPtr := getopt.BoolLong("show-rules", 'l', "view all available rules")
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

    pL := clt.NewLoadingMessage("Loading...", clt.Dots, 100*time.Millisecond)
	pL.Start()
	time.Sleep(4 * time.Second)
	pL.Success()


    p := clt.NewIncrementalProgressBar(10, "Doing work")

	ch := make(chan int, 1)
	var wg sync.WaitGroup

	// start 3 go routines to do some work.  Pass them a copy of the progress bar so they can
	// call increment after each task is done
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(ch chan int, p *clt.Progress) {
			defer wg.Done()
			for range ch {
				time.Sleep(time.Duration(rand.Intn(1000) * 1000000))
				p.Increment()
			}
		}(ch, p)
	}

	// start the bar then pass it some work to do
	p.Start()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// wait until all the work is done
	wg.Done()

	// call Success to close the progress channel and update to 100%
	//p.Success()
	p.Fail()
}