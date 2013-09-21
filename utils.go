package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func createTempDir(prefix string) string {
	path, err := ioutil.TempDir("/tmp", prefix)
	if err != nil {
		log.Fatalln("Unable to create build directory.")
	}
	return path
}

func createDir(path string) {
	err := os.Mkdir(path, 0700)
	if err != nil {
		log.Fatalln("Unable to create directory.", path)
	}
}

func getSrcDir() string {
	_, name, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Cannot find source directory.")
	}
	return filepath.Dir(name)
}
