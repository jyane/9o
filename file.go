package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic("file read error.")
	}
	defer f.Close()
	l, err := ioutil.ReadAll(f)
	if err != nil {
		panic("file read error.")
	}
	return string(l)
}

func Write(path string, asm string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic("file write error")
	}
	defer f.Close()
	fmt.Fprintln(f, asm)
}
