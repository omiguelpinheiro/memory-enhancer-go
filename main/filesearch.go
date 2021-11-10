package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/icza/backscanner"
)

func ReverseSearch(scanner *backscanner.Scanner, search []byte) (position int, line string, err error) {
	for {
		line, pos, err := scanner.LineBytes()
		if err != nil {
			if err == io.EOF {
				return -1, "", fmt.Errorf((fmt.Sprintf("%q is not found in file.\n", search)))
			} else {
				return -1, "", err
			}
		}
		if bytes.Contains(line, search) {
			return pos, string(line), nil
		}
	}
}