package jquery

import (
	"io/ioutil"

	"github.com/gocaveman/caveman/uifiles/uiregistry"
	"github.com/gocaveman/caveman/webutil"
)

//go:generate go run assets_generate.go

func init() {

	f, err := assets.Open("jquery.js")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	bds := webutil.NewBytesDataSource(b, "jquery.js", st.ModTime())
	uiregistry.MustRegister("js:github.com/gocaveman-libs/jquery", nil, bds)

}
