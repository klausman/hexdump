// Copyright 2019 Tobias Klausmann
// Licensed under the Apache License v2. See LICENSE for details

package hexdump

import (
	"fmt"
)

func ExampleDump() {
	fmt.Println(Dump([]byte("Hello world!\n\tAnd sundry other stuff\n"), 0, 4))
	// Output: 0000  48 65 6c 6c 6f 20 77 6f 72 6c 64 21 0a 09 41 6e Hello world!..An
	// 0010  64 20 73 75 6e 64 72 79 20 6f 74 68 65 72 20 73 d sundry other s
	// 0020  74 75 66 66 0a                                  tuff.
}
