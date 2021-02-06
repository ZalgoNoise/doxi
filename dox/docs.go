package dox

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// Docs struct will define a directory for documentation, containing
// reference to the actual path to the folder, the project name, and
// a list of File objects for each created document
type Docs struct {
	Path    string `json:"path"`
	Proj    string `json:"project"`
	Set     string `json:"doc_set"`
	Content []File `json:"files"`
}

// NewDocs function will create a new instance of Docs
func NewDocs() *Docs {
	d := &Docs{}
	return d
}

// Dir method will scan the Dir object and replay its contents, but inside
// the /docs directory. Instead of creating the files straight away, this
// part of the process will simply gather the metadata (references for source)
// and building the new Files objects (and linking them to the original ones)
// in order to branch off of this object with any required actions
func (d *Docs) Dir(i *Dir) *Dir {

	dir := os.Getenv("PWD")

	os.Chdir(d.Path)

	defer os.Chdir(dir)

	path := d.Path + "/"

	for idx, v := range i.Files {

		srcRelPath := d.Path + "/src" + v.RelPath.Get()
		srcParPath := path + "src/" + v.Parent.Get()
		pkgRelPath := d.Path + "/pkg" + v.RelPath.Get()
		pkgParPath := path + "pkg/" + v.Parent.Get()

		if v.Type == "dir" {
			if strings.Contains(srcRelPath, "/docs") {
				continue
			}
			if strings.Contains(srcRelPath, "/bazel-") {
				continue
			}
			if strings.Contains(srcRelPath, "/testdata") {
				continue
			}
			os.Mkdir(srcRelPath, 0755)
			os.Mkdir(pkgRelPath, 0755)
		}

		if v.Type == "file" && v.Ext == ".go" {
			if _, err := os.Stat(srcParPath); os.IsNotExist(err) {
				os.Mkdir(srcParPath, 0755)
				os.Mkdir(pkgParPath, 0755)
			}
			newSrc, err := os.Create(srcRelPath + ".md")
			newPkg, err := os.Create(pkgRelPath + ".md")
			defer newSrc.Close()
			defer newPkg.Close()

			srcF := NewFile()
			pkgF := NewFile()

			srcF.Init(d.Path, srcRelPath+".md")
			i.Files[idx].Link(*srcF)
			srcF.Link(v)

			pkgF.Init(d.Path, pkgRelPath+".md")
			i.Files[idx].Link(*pkgF)
			pkgF.Link(v)

			d.Content = append(d.Content, *srcF)

			if err != nil {
				panic(err)
			}

		}

	}

	return i

}

// Source method will create a new SourceCode object for each Go file,
// read its contents, get a symlink to the original object, and finally
// generate the Markdown file based on this content
func (d *Docs) Source(i *Dir) *Dir {

	for _, v := range i.Files {

		if v.Ext == ".go" {
			src := NewSourceCode()

			src.Read(v.Path.Get())
			src.Link(&v)
			//src.GenSrc(v.Source)
			src.GenPkg(v.Source)
		}
	}

	return i
}

func run(args ...string) ([]byte, error) {

	cmdPath, err := exec.LookPath(args[0])
	if err != nil {
		return nil, err
	}

	cmdFlags := args[1:]

	c := exec.Command(cmdPath, cmdFlags...)
	var outb, errb bytes.Buffer
	c.Stdout = &outb
	c.Stderr = &errb
	e := c.Run()
	if e != nil {
		return nil, e
	}

	return outb.Bytes(), nil

}
