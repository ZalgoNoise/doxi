package dox

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// Dir struct is the placeholder for the current path which is perceived
// as the go project path, and will initiate the process of documenting
// Go files
type Dir struct {
	Path  string
	Files []File
	Docs  Docs
}

// New function will create and initialize a new Dir object, and return it
func New() *Dir {
	d := &Dir{}

	d.Init()
	return d
}

// Init method will initialize a new Dir object with the current path
// and scan the existing files in the directory, creating File objects
// along the way
func (d *Dir) Init() *Dir {
	d.Path = os.Getenv("PWD")
	d.Check("/docs")
	d.Check("/docs/src")
	d.Check("/docs/pkg")

	var files []string

	r := d.Path
	err := filepath.Walk(
		r,
		func(p string, i os.FileInfo, err error) error {
			files = append(files, p)
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
	for _, v := range files {
		if v != d.Path {
			f := NewFile()
			f.Init(d.Path, v)
			d.Files = append(d.Files, *f)
		}
	}

	return d
}

// Check method will look if the /docs directory exist and create it
// if it doesn't
func (d *Dir) Check(s string) {
	path := d.Path + s

	if _, err := os.Stat(path); os.IsNotExist(err) {

		e := os.Mkdir(path, 0755)
		if e != nil {
			panic(e)
		}
	}
	if _, err := os.Stat(path); os.IsExist(err) {
		os.RemoveAll(path)
		e := os.Mkdir(path, 0755)
		if e != nil {
			panic(e)
		}
	}

}

// Run method will execute Doxi for this project, and generate Markdown
// documentation for your Go files
func (d *Dir) Run() *Dir {
	docs := NewDocs()
	d.Docs = *docs

	path := d.Path + "/docs"

	docs.Path = path
	proj := strings.Split(d.Path, "/")
	docs.Proj = proj[(len(proj) - 1)]

	docs.Dir(d)

	docs.Source(d)

	return d

}

// Done method will provide output on the execution
func (d *Dir) Done() []byte {

	json, err := json.Marshal(d.Docs)
	if err != nil {
		panic(err)
	}
	return json
}
