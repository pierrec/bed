//+build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"path/filepath"
	"strings"
)

func genUnpack64Table() []unpack64Entry {
	// Skip first and last entry as they are trivial.
	table := make([]unpack64Entry, 254)

	for i := range table {
		entry := &table[i]
		i++
		i := bits.Reverse8(uint8(i))
		entry.num = bits.OnesCount8(i)
		shift := bits.TrailingZeros8(i) / 8
		for j := 0; i > 0; {
			if i&1 > 0 {
				entry.shifts[j] = 8 * shift
				j++
			}
			i >>= 1
			shift++
		}
	}

	return table
}

type unpack64Entry struct {
	num    int // number of non zero bytes
	shifts [8]int
}

func printUnpack64Table(w io.Writer, t []unpack64Entry) (err error) {
	_, err = fmt.Fprint(w, "var unpack64Table = [...]unpack64Entry{")
	if err != nil {
		return
	}
	for i, e := range t {
		if i%8 == 0 {
			_, err = fmt.Fprintf(w, "\n\t")
			if err != nil {
				return
			}
		}

		// Drop the last shift as it is only set for 255.
		_, err = fmt.Fprintf(w, "{%d,%d,%d,%d,%d,%d,%d,%d}, ", e.num,
			e.shifts[0], e.shifts[1], e.shifts[2], e.shifts[3], e.shifts[4], e.shifts[5], e.shifts[6])
		if err != nil {
			return
		}
	}
	_, err = fmt.Fprint(w, "\n}\n")
	return
}

func main() {
	out, err := os.Create("uint_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	var header = `// Code generated by "%s". DO NOT EDIT.

package serializer

`
	cmd := append([]string{filepath.Base(os.Args[0])}, os.Args[1:]...)
	if _, err := fmt.Fprintf(out, header, strings.Join(cmd, " ")); err != nil {
		log.Fatal(err)
	}

	if err := printUnpack64Table(out, genUnpack64Table()); err != nil {
		log.Fatal(err)
	}
}
