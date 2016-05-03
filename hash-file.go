package main

import (
	"flag"

	hashfile "github.com/pschlump/hash-file/lib"
)

var Output = flag.String("output", "", "Output - defaults to standard output") // 0
func init() {
	flag.StringVar(Output, "o", "", "Output - defaults to standard output") // 0
}

func main() {

	flag.Parse()
	fns := flag.Args()

	cfg := &hashfile.HashLibCfg{}

	hashfile.HashFiles(cfg, *Output, fns...)

}
