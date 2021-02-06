package dox

import (
	"os"
	"strings"
)

// File struct will describe a file object, within the context of a
// project. It is supposed to hold all required metadata to build
// documentation with Doxi.
type File struct {
	Proj    Proj    `json:"project"`
	Path    Path    `json:"abs_path"`
	RelPath RelPath `json:"rel_path"`
	Parent  Parent  `json:"parent_dir"`
	Name    Name    `json:"file_name"`
	Ext     Ext     `json:"extension"`
	Type    Type    `json:"type"`
	Source  *File   `json:"mem_ref"`
}

// Proj type will represent the project name (parent folder to current dir)
type Proj string

// Path is the absolute path to the project's directory
type Path string

// RelPath is the relative path within the project's directory
type RelPath string

// Parent is the parent folder for this file
type Parent string

// Name is the name of the file
type Name string

// Ext is the file's extension
type Ext string

// Type will describe if it's a file or folder (folders are files in Unix)
type Type string

// NewFile function will create a new instance of File
func NewFile() *File {
	f := &File{}
	return f
}

// Init method will initialize a new file based on its (base) path and
// the actual file path
func (f *File) Init(p, s string) *File {
	f.Path.Set(f, s)

	proj := f.Proj.Gen(f, p)
	f.Proj.Set(f, proj)

	file := strings.Split(s, p)
	if len(file) > 1 {
		f.RelPath.Set(f, file[1])
	}

	parent := strings.Split(file[1], "/")
	f.Parent.Set(f, parent[(len(parent)-2)])

	f.Name.Set(f, parent[(len(parent)-1)])

	ext := strings.Split(file[1], ".")
	if len(ext) > 1 {
		f.Ext.Set(f, "."+ext[(len(ext)-1)])
		f.Type.Set(f, "file")
	} else if _, err := os.Stat(s); !os.IsNotExist(err) {
		f.Type.Set(f, "dir")
	}

	return f
}

// Unite method will create mutual symlinks between File objects
// (memory-wise)
func (f *File) Unite(i File) File {
	f.Source = &i
	i.Source = f
	return i

}

// Fetch method will return the source / origin / link of a requested file
func (f *File) Fetch() File {
	return *f.Source
}

// Link method will create a memory symlink with both File objects
func (f *File) Link(i File) {
	f.Source = &i
}

// Gen method - Proj - Generate the reference project name based on the
// current file's relative path. If the object derives from a Docs object,
// jump one folder up; otherwise it's the last "entry" in the path
func (f *Proj) Gen(i *File, s string) string {

	path := strings.Split(s, "/")

	if path[(len(path)-1)] == "docs" {
		return path[(len(path) - 2)]
	}
	return path[(len(path) - 1)]
}

// Set method - Proj - defines this type with an input File and string
func (f *Proj) Set(i *File, s string) {
	new := Proj(s)
	i.Proj = new
}

// Set method - Path - defines this type with an input File and string
func (f *Path) Set(i *File, s string) {
	new := Path(s)
	i.Path = new
}

// Set method - RelPath - defines this type with an input File and string
func (f *RelPath) Set(i *File, s string) {
	new := RelPath(s)
	i.RelPath = new
}

// Set method - Parent - defines this type with an input File and string
func (f *Parent) Set(i *File, s string) {
	new := Parent(s)
	i.Parent = new
}

// Set method - Name - defines this type with an input File and string
func (f *Name) Set(i *File, s string) {
	new := Name(s)
	i.Name = new
}

// Set method - Ext - defines this type with an input File and string
func (f *Ext) Set(i *File, s string) {
	new := Ext(s)
	i.Ext = new
}

// Set method - Type - defines this type with an input File and string
func (f *Type) Set(i *File, s string) {
	new := Type(s)
	i.Type = new
}

// Get method - Proj - retrieves the string value for this type
func (f *Proj) Get() string {
	return string(*f)
}

// Get method - Path - retrieves the string value for this type
func (f *Path) Get() string {
	return string(*f)
}

// Get method - RelPath - retrieves the string value for this type
func (f *RelPath) Get() string {
	return string(*f)
}

// Get method - Parent - retrieves the string value for this type
func (f *Parent) Get() string {
	return string(*f)
}

// Get method - Name - retrieves the string value for this type
func (f *Name) Get() string {
	return string(*f)
}

// Get method - Ext - retrieves the string value for this type
func (f *Ext) Get() string {
	return string(*f)
}

// Get method - Type - retrieves the string value for this type
func (f *Type) Get() string {
	return string(*f)
}
