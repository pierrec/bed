package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	bi, _ := debug.ReadBuildInfo()
	fmt.Fprintf(os.Stdout, "%#v", bi)
}
