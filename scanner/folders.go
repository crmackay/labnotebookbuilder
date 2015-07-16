/*package mylabnotebook

import (
	//"fmt"
	"time"
)

type folder struct {
}

type experiment struct {
}

type project struct {
}

type file struct {
	name    string
	path    string
	ext     string
	content []byte
}

type expEntry struct {
	file
	date   time.Time
	title  string
	parent experiment
}

func newFile() {}
*/

package scanner

import (
	"fmt"
	"path/filepath"
    "os"
)

type File struct {
	name    string
	kind    string
	absPath string
	parent  string
}

func NewFile(path string) {
    name :=
    kind :=
    splitPath := Split(path, os.PathSeparator)

    for i, elem := range splitPath {

    }
}

type Direc struct {
	splitPath    []string
	name     string
	parent   string
	children string
	files    []string
}


// create a map of 
