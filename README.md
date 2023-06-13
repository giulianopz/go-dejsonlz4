# go-dejsonlz4

Decompress Firefox bookmark files with `.jsonlz4` extension, a plain `lz4` format but with custom header (so-called [mozLz4](http://justsolve.archiveteam.org/wiki/Mozilla_LZ4)).

## Usage

Run it from the terminal:
```bash
# build from source 
:~$ go build -o go-dejsonlz4
# or, install from the GitHub repo
:~$ go install github.com/giulianopz/go-dejsonlz4@latest
:~$ go-dejsonlz4 -h
Usage: go-dejsonlz4 [-h] IN_FILE [OUT_FILE]
Example: go-dejsonlz4 ~/.mozilla/firefox/$PROFILE_NAME/bookmarkbackups/$FILE_NAME.jsonlz4
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

- the original C code to perform decompression used by [Mozilla](https://hg.mozilla.org/mozilla-central/file/c3f5e6079284/mfbt/lz4.c)
- [LZ4 explained](http://fastcompression.blogspot.com/2011/05/lz4-explained.html)
- [Go porting of the C reference version](https://lz4.github.io/lz4/)
