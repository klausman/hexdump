// Copyright 2019 Tobias Klausmann
// Licensed under the Apache License v2. See LICENSE for details

// Package hexdump provides a simple hexdumper for binary data.
package hexdump

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// Dump returns a string that represents the supplied data in the classic
// hexdump format. The offset parameter sets the starting offset of the
// leftmost column. The last parameter sets the width of the offset column.
// Data is always dumped in sixteen space-separated columns,
// with the ASCII-printable at the right.
func Dump(b []byte, offset, width int) string {
	i := 0
	l := len(b)
	r := []string{}
	for i < l {
		d := b[i:intMin(i+16, l)]
		r = append(r, fmt.Sprintf("%0*x  %s %s\n", width, offset+i, hexCol(d, 16), asciionly(d)))
		i += 16
	}
	return strings.Join(r, "")
}

func hexCol(d []byte, pad int) string {
	var r []string
	hs := hex.EncodeToString(d)
	for x := 0; x < len(hs); x += 2 {
		r = append(r, hs[x:x+2])
	}
	if (len(d) % pad) != 0 {
		r = append(r, mulStr("  ", pad-(len(d)%pad))...)
	}
	return strings.Join(r, " ")
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asciionly(d []byte) []byte {
	var r []byte
	for _, b := range d {
		if (byte(31) < b) && (b < byte(127)) {
			r = append(r, b)
		} else {
			r = append(r, 46)
		}
	}
	return r
}

func mulStr(s string, c int) []string {
	var r []string
	for i := 0; i < c; i++ {
		r = append(r, s)
	}
	return r
}
