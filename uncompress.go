package main

import (
	"fmt"
	"reflect"

	"github.com/pierrec/lz4"
)

// i.e. "mozLz40\0"
var magicHeader [8]int = [8]int{109, 111, 122, 76, 122, 52, 48, 0}

// the 4 bytes after the header store the size of the original (uncrompressed) file
const decompSize int = 4

const magicSize = len(magicHeader)

// Uncompress decompresses input data from a Firefox bookmark file with .jsonlz4 extension
func Uncompress(inputData []byte) ([]byte, error) {

	var inputSize int = len(inputData)

	// read and check magic header
	if inputSize < magicSize+decompSize || reflect.DeepEqual(magicHeader, inputData[:magicSize]) {
		return nil, fmt.Errorf("unsupported file format")
	}

	// decode size of decompressed output
	// for standard lz4, which does not include such info, use a heuristic func as suggested by Mark Adler: https://stackoverflow.com/a/25755758/9109880
	var outputSizeUpperBound int
	for i := magicSize; i < magicSize+decompSize; i++ {
		outputSizeUpperBound += (int)(inputData[i]) << (8 * (i - magicSize))
	}

	// read and uncompress payload
	outputData := make([]byte, outputSizeUpperBound)
	_, err := lz4.UncompressBlock(inputData[magicSize+decompSize:], outputData)
	if err != nil {
		return nil, err
	}

	return outputData, nil
}
