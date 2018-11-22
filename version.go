package main

import (
	"fmt"
	"runtime"
)

var (
	// Variables defined by the Makefile and passed in with ldflags
	Version   = "devel"
	CommitSHA = "unknown"
	BuildDate = "unknown"
)

func printVersion() {
	fmt.Printf(`aws-power-toggle:
  version     : %s
  build date  : %s
  git hash    : %s
  go version  : %s
  go compiler : %s
  platform    : %s/%s
`, Version, BuildDate, CommitSHA, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
}

func getVersionResponse() string {
	return fmt.Sprintf(`{"version":"%s","git_hash":"%s","build_date":"%s"}`, Version, CommitSHA, BuildDate)
}
