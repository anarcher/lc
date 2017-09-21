package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-logfmt/logfmt"
)

func main() {
	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	d := logfmt.NewDecoder(os.Stdin)

	for d.ScanRecord() {
		for d.ScanKeyval() {
			key := string(d.Key())
			value := string(d.Value())

			switch value {
			case "":
				fmt.Printf("%s ", key)
			case "err":
				fmt.Print("%s=%s ", red(key), value)
			default:
				fmt.Printf("%s=%s ", cyan(key), value)
			}

		}
		fmt.Println()
	}

	if d.Err() != nil {
		fmt.Println("%s=", red("lc.err"), d.Err())
		os.Exit(1)
	}
}
