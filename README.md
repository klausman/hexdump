# Hexdump

**NOTE**: As of 2024-06-23 the Go import url of this repo has changed to:
`pkg.i-no.de/pkg/hexdump`. Issues, merge requests etc should be filed at its new
source location, https://codeberg.org/klausman/hexdump

[![GoDoc](https://godoc.org/pkg.i-no.de/pkg/hexdump?status.svg)](https://godoc.org/pkg.i-no.de/pkg/hexdump)

A simple hexdumper in Go. Produces output like this:

```
0000  48 65 6c 6c 6f 20 77 6f 72 6c 64 21 0a 09 41 6e Hello world!..An
0010  64 20 73 75 6e 64 72 79 20 6f 74 68 65 72 20 73 d sundry other s
0020  74 75 66 66 0a                                  tuff.
```

Starting offset and width of the offset column are configurable.


