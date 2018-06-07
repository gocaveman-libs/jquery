// +build ignore

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const libname = "jquery"

var files = []string{
	"jquery.js",
}

func main() {

	var srcbuf bytes.Buffer
	fmt.Fprintf(&srcbuf, `package %s`+"\n", libname)
	fmt.Fprintf(&srcbuf, `func getEntries() (string, []entry) {`+"\n")
	fmt.Fprintf(&srcbuf, `    var entries []entry`+"\n")

	for _, file := range files {
		err := addFile(&srcbuf, file)
		if err != nil {
			panic(err)
		}
	}

	fmt.Fprintf(&srcbuf, `    return %q, entries`+"\n", libname)
	fmt.Fprintf(&srcbuf, `}`+"\n")

	err := ioutil.WriteFile("webresource-data.go", srcbuf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

}

func addFile(w io.Writer, name string) error {

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	_, err = gw.Write(b)
	if err != nil {
		return err
	}
	err = gw.Close()
	if err != nil {
		return err
	}

	fmt.Fprintf(w, `// compressed size: %d`+"\n", buf.Len())
	fmt.Fprintf(w, `entries = append(entries, entry{t:%d, gb:[]byte(%q)})`+"\n", fi.ModTime().Unix(), buf.String())

	return nil
}
