package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var dir string
	var err error

	dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		fmt.Printf("%s  ", files[i].Name())
	}
}
