package dox

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var newLine = []byte(`
`)

// SourceCode struct will contain one instance of a source code file
// with all required parameters to build one, plus a pointer to the
// original file
type SourceCode struct {
	Content []string
	Buffer  []byte
	Lines   int
	Proj    string
	File    string
	Path    string
	RelPath string
	Ref     *File
}

// NewSourceCode function will create a new instance of SourceCode
func NewSourceCode() *SourceCode {
	return &SourceCode{}
}

// Read method will define the SourceCode object as per the file in the
// provided path
func (s *SourceCode) Read(path string) *SourceCode {
	var content []string

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	s.Content = content
	s.Lines = len(content)

	return s
}

// Link method will point a SourceCode object to its original File
func (s *SourceCode) Link(f *File) *SourceCode {
	s.Proj = f.Proj.Get()
	s.Ref = f
	s.RelPath = f.RelPath.Get()
	return s
}

// SetPath method will point the files to the set /docs/* folder
func (s *SourceCode) SetPath(p string) string {
	paths := strings.Split(s.Ref.Path.Get(), s.Ref.RelPath.Get())
	srcPath := paths[0] + "/docs/" + p + s.Ref.RelPath.Get() + ".md"

	return srcPath
}

// GenSrc method will create the /docs/src content for the SourceCode
// generated document
func (s *SourceCode) GenSrc(f *File) *SourceCode {
	s.File = f.Name.Get()
	s.Path = s.SetPath("src")

	header := s.GenHeader("src")

	code := s.GenCode()

	var buf []byte
	s.Buffer = byteJoin(buf, header, code)

	err := ioutil.WriteFile(s.Path, s.Buffer, 0644)
	if err != nil {
		panic(err)
	}

	return s
}

// GenPkg method will create the /docs/pkg content for the SourceCode
// generated document
func (s *SourceCode) GenPkg(f *File) *SourceCode {
	s.File = f.Name.Get()
	s.Path = s.SetPath("pkg")

	//header := s.GenHeader("pkg")

	return s
}

func (s *SourceCode) GenHeader(docs string) []byte {

	hPath := strings.Split(s.RelPath, "/")
	hPath[0] = docs

	var headerURL []string

	for i := 0; i < len(hPath); i++ {
		headerURL = append(headerURL, mdRelURL(hPath[i], (len(hPath)-1-i)))
	}

	var hURL string

	for _, v := range headerURL {
		hURL += "/" + v
	}

	return []byte("## " + hURL + "\n\n")

}

func (s *SourceCode) GenCode() []byte {
	var gen []string
	var buf []byte

	codeFmt := []byte("<pre class=" + `"code highlight"` + "><code>\n")

	buf = byteJoin(buf, codeFmt)
	for idx, v := range s.Content {
		line := strconv.Itoa(idx + 1)
		entry := `<span id="L` + line + `" class="line" lang="go"><a href="#L` + line + `">` + line + `</a>	</span><span class="comment">` + v + `</span>`

		gen = append(gen, entry)
		buf = byteJoin(buf, []byte(entry), newLine)
	}

	footer := []byte("</code></pre>\n\n_____")

	buf = byteJoin(buf, footer)

	s.Content = gen
	return buf

}

func byteJoin(input ...[]byte) []byte {
	var empty []byte
	array := bytes.Join(input, empty)
	return array
}

func mdURL(s, u string) string {
	return "[" + s + "](" + u + ")"
}

func mdRelURL(s string, i int) string {

	if i <= 0 {
		return mdURL(s, "./"+s+".md")

	}
	var relPath string

	for idx := 1; idx <= i; idx++ {
		relPath += "../"
	}
	return mdURL(s, relPath+s+"/index.md")

}
