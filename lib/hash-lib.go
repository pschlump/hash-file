package hashlib

import (
	"fmt"
	"os"

	"github.com/pschlump/Go-FTL/server/sizlib"
	"github.com/pschlump/HashStrings"
)

// "crypto/sha256"

type HashLibCfg struct {
	// add line numbering
	// add de-tabbing
	// others...
}

func HashFiles(cfg *HashLibCfg, output string, files ...string) {
	var err error

	of := os.Stdout
	if output != "" {
		of, err = sizlib.Fopen(output, "w")
		if err != nil {
			fmt.Fprintf(os.Stderr, "go-cat: Error: Unable to open output file %s, %s\n", output, err)
			os.Exit(1)
		}
	}

	for _, vv := range files {
		ifp, err := sizlib.Fopen(vv, "r")
		if err != nil {
			fmt.Fprintf(os.Stderr, "hash-file: Error: Unable to open file %s for read, %s\n", vv, err)
			continue
		}
		buf := make([]byte, 10240)
		n1, err := ifp.Read(buf) // n1 is # of bytes read
		if err != nil {
			fmt.Fprintf(os.Stderr, "hash-file: Error: Unable to read file %s, %s\n", vv, err)
		}
		// fmt.Printf("n1=%d buf >%s<\n", n1, buf)
		fmt.Printf("n1=%d\n", n1)
		ifp.Close()
		// hash of data + timestamp + length
		etag := HashStrings.Sha256(string(buf[0:n1]))
		// printout hash
		fmt.Fprintf(of, "%40s %s\n", etag, vv)
	}

}
