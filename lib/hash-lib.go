package hashlib

import (
	"fmt"
	"os"
	"time"

	"github.com/pschlump/Go-FTL/server/sizlib"
	"github.com/pschlump/HashStrings"
)

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
		hh, err := HashFile(cfg, vv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "hash-file: %s\n", err)
		} else {
			fmt.Fprintf(of, "%40s %s\n", hh, vv)
		}
	}
}

func HashFile(cfg *HashLibCfg, fn string) (h string, err error) {
	fInfo, err := os.Stat(fn)
	if err != nil {
		err = fmt.Errorf("Error: Unable to stat file %s, %s\n", fn, err)
		return
	}
	size := fInfo.Size()
	mTime := fInfo.ModTime()
	sMTime := mTime.Format(time.RFC3339Nano)
	ifp, err := sizlib.Fopen(fn, "r")
	if err != nil {
		err = fmt.Errorf("Error: Unable to open file %s for read, %s\n", fn, err)
		return
	}
	defer ifp.Close()
	buf := make([]byte, 10240+35+15)
	buf1 := fmt.Sprintf("%35s%15d", sMTime, size)
	// fmt.Printf("buf1 [%s]\n", buf1)
	copy(buf, buf1)
	// fmt.Printf("Before: %s\n", buf[0:50])
	n1, err := ifp.Read(buf[35+15:]) // n1 is # of bytes read
	if err != nil {
		err = fmt.Errorf("Error: Unable to read file %s, %s\n", fn, err)
		return
	}
	// fmt.Printf("n1=%d buf >%s<\n", n1, buf)
	// fmt.Printf("n1=%d\n", n1)
	// hash of data + timestamp + length
	h = HashStrings.Sha256(string(buf[0 : n1+35+15]))
	// printout hash
	return
}

// NewETag, err := hashlib.HashData(newdata)
func HashData(data []byte) (h string, err error) {
	h = HashStrings.Sha256(string(data))
	return
}

// NewETag, err := hashlib.HashFile(nil, name) // xyzzy - shoud have a HashFileWithInfo ( nil, name, size, modtime )

/*

// if found2, fileInfo2 := lib.ExistsGetFileInfo(index); found2 {
// fileInfo.ModTime().Format(time.RFC3339Nano)

func ExistsGetFileInfo(name string) (bool, os.FileInfo) {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	return true, fi
}


sizeFunc := func() (int64, error) { return fileInfo.Size(), nil }

func ServeContent(w http.ResponseWriter, req *http.Request, name string, modtime time.Time, content io.ReadSeeker) {
	sizeFunc := func() (int64, error) {
		size, err := content.Seek(0, os.SEEK_END)
		if err != nil {
			return 0, errSeeker
		}
		_, err = content.Seek(0, os.SEEK_SET)
		if err != nil {
			return 0, errSeeker
		}
		return size, nil
	}
	serveContent(w, req, name, modtime, sizeFunc, content)
}

*/
