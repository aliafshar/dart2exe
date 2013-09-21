package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

type Environ struct {
	buildName        string
	buildPkg         string
	buildPrefix      string
	buildDir         string
	buildSrcDir      string
	buildBinDir      string
	buildBinExec     string
	buildTmpDir      string
	buildPkgDir      string
	buildFuncPrefix  string
	buildChunkSize   int
	srcDir           string
	dartVmPath       string
	dartPubPath      string
	dartPkgPath      string
	dartPkgName      string
	mainTemplateName string
	mainTemplatePath string
	zipName          string
	zipPath          string
	cwd              string
}

func NewEnviron() *Environ {
	flag.Parse()
	e := &Environ{
		buildName:      "dart2exe",
		buildChunkSize: 1024 * 32,
		dartVmPath:     "/usr/local/dart/dart-sdk/bin/dart",
		dartPubPath:    "/usr/local/dart/dart-sdk/bin/pub",
	}
	e.srcDir = getSrcDir()
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	e.cwd = cwd
	if len(flag.Args()) == 0 {
		e.dartPkgPath = filepath.Join(e.srcDir, "dart_test_app")
	} else {
		e.dartPkgPath = flag.Arg(0)
	}
	e.dartPkgName = filepath.Base(e.dartPkgPath)
	if len(flag.Args()) > 1 {
		e.buildBinExec = flag.Arg(1)
	} else {
		e.buildBinExec = filepath.Join(e.cwd, e.dartPkgName+"_2exe")
	}
	e.buildPkg = e.buildName + "_bootstrap"
	e.buildPrefix = e.buildName + "-build-"
	e.buildDir = createTempDir(e.buildPrefix)
	e.buildSrcDir = filepath.Join(e.buildDir, "src")
	e.buildBinDir = filepath.Join(e.buildDir, "bin")
	e.buildTmpDir = filepath.Join(e.buildDir, "tmp")
	e.buildPkgDir = filepath.Join(e.buildSrcDir, e.buildPkg)
	e.buildFuncPrefix = e.buildPkg
	e.mainTemplateName = "src_main.go.txt"
	e.mainTemplatePath = e.srcPath(e.mainTemplateName)
	e.zipName = e.buildName + ".tar"
	e.zipPath = filepath.Join(e.buildTmpDir, e.zipName)
	createDir(e.buildSrcDir)
	createDir(e.buildBinDir)
	createDir(e.buildTmpDir)
	createDir(e.buildPkgDir)
	return e
}

func (e *Environ) buildPkgPath(name string) string {
	return filepath.Join(e.buildPkgDir, name)
}

func (e *Environ) buildBinPath(name string) string {
	return filepath.Join(e.buildBinDir, name)
}

func (e *Environ) srcPath(name string) string {
	return filepath.Join(e.srcDir, name)
}
