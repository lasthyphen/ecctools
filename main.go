package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/lasthyphen/ecctools/cmd"
	"github.com/lasthyphen/ecctools/pkg/version"
)

func main() {
	defer panicHandler()
	cmd.Execute()
}

func panicHandler() {
	if panicPayload := recover(); panicPayload != nil {
		stack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintln(os.Stderr, "            Fatal error. Sorry! You found a bug.")
		fmt.Fprintln(os.Stderr, "    Please copy all of this info into an issue at")
		fmt.Fprintln(os.Stderr, "     https://github.com/lasthyphen/ecctools")
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintf(os.Stderr, "Version:           %s\n", version.Version)
		fmt.Fprintf(os.Stderr, "Build Date:        %s\n", version.BuildDate)
		fmt.Fprintf(os.Stderr, "Git Commit:        %s\n", version.GitCommit)
		fmt.Fprintf(os.Stderr, "Go Version:        %s\n", version.GoVersion)
		fmt.Fprintf(os.Stderr, "OS / Arch:         %s\n", version.OsArch)
		fmt.Fprintf(os.Stderr, "Panic:             %s\n\n", panicPayload)
		fmt.Fprintln(os.Stderr, stack)
		os.Exit(1)
	}
}
