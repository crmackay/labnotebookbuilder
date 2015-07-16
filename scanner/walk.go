package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Walk Func ..
func WalkFunc(path string, f os.FileInfo, err error) error {
	fmt.Println("hello")
	return nil
}

func main() {

	path := "/User/christophermackay/lab/notebook"
	fmt.Println("test")

	filepath.Walk(path, WalkFunc)
}
