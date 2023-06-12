# go-dejsonlz4

Decompress Firefox bookmark files with `.jsonlz4` extension, a plain `lz4` format but with custom header (so-called [mozLz4](http://justsolve.archiveteam.org/wiki/Mozilla_LZ4)).


## Usage

Run it from the terminal:
```bash
:~$ go build -o go-dejsonlz4
:~$ ./go-dejsonlz4 -h
Usage: go-dejsonlz4 [-h] IN_FILE [OUT_FILE]
Example: go-dejsonlz4 ~/.mozilla/firefox/aks8v8c0.default-release/bookmarkbackups/bookmarks-2023-03-01_1011_OItiw5WByHsdl6u-lQ08mQ==.jsonlz4
Decompress Firefox bookmark files with .jsonlz4 extension from IN_FILE to OUT_FILE:
        * -h, display help message and exit,
        * IN_FILE='-', uncompress from standard input,
        * OUT_FILE='-' or missing, uncompress to standard output.
```

Or use it in your go module:
```go
package main

import (
	"fmt"
	"os"

	dejsonlz4 "github.com/giulianopz/go-dejsonlz4"
)

func main() {

	inputData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	outputData, err := dejsonlz4.Uncompress(inputData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(outputData))
}
```

## References

The original C code to perform decompression can be found in the Mozilla Mercurial [repository](https://hg.mozilla.org/mozilla-central/file/c3f5e6079284/mfbt/lz4.c).
