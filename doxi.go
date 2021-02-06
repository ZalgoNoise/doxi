// Doxi is a Go source code documentation generator, which generates static
// Markdown / HTML files for you to host. The point is to generate a set of
// documents similar to the Golang source code documentation, as seen with
// `godoc` and `go doc`, however having a static Markdown / HTML render,
// with relative hyperlinks instead of absolute ones.
//
// You can easily document your project by running the `doxi` binary, or by
// embedding simple logic as below into your existing project, as a closer,
// or "onDone" type of function:
//
//     package main
//
//     import (
//	       "github.com/ZalgoNoise/doxi/dox"
//     )
//
//     func main() {
//         // your code goes here
//         dox := dox.New()
//         dox.Run()
//     }
//
//
package main

import (
	"github.com/ZalgoNoise/doxi/dox"
)

func main() {
	dox := dox.New()
	dox.Run()

}
