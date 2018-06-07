package jquery

import (
	"time"

	"github.com/gocaveman/webresource" // FIXME: will change to correct import, this is just for the prototype
)

//go:generate go run webresource-generate.go
//go:generate go fmt webresource-data.go

type entry struct {
	n  string
	t  int64
	gb []byte
}

func ModulePROTO() webresource.Module {
	name, entries := getEntries()
	fs := webresource.NewFileSet(name /*, ...dependencies go here... */)
	for _, e := range entries {
		fs = fs.WriteGzipFile("/"+e.n, 0644, time.Unix(e.t), e.gb)
	}
	return fs
}
