package main

import (
	_ "flag"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Chunk struct {
	env   *Environ
	buf   []byte
	index int
}

func (c *Chunk) createData(chunks chan *Chunk) {
	cmd := exec.Command("go-bindata", "-func="+c.FuncName(), "-nomemcopy")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("Failed to get bindata input pipe.", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("Failed to get bindata output pipe.", err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatalln("Failed to start bindata.", err)
	}
	_, err = stdin.Write(c.buf)
	if err != nil {
		log.Fatalln("Failed to write bindata input pipe.", err)
	}
	err = stdin.Close()
	if err != nil {
		log.Fatalln("Failed to close bindata input pipe.", err)
	}
	f, err := os.Create(c.outPath())
	if err != nil {
		log.Fatalln("Failed to create output file.", err)
	}
	_, err = io.Copy(f, stdout)
	if err != nil {
		log.Fatalln("Failed to write output file.", err)
	}
	f.Close()
	err = cmd.Wait()
	if err != nil {
		log.Fatalln("Failed to wait for bindata.", err)
	}
	chunks <- c
}

func (c *Chunk) FuncName() string {
	return c.env.buildFuncPrefix + strconv.Itoa(c.index)
}

func (c *Chunk) outPath() string {
	return c.env.buildPkgPath(c.FuncName() + ".go")
}

type Builder struct {
	env    *Environ
	chunks []*Chunk
}

func NewBuilder(env *Environ) *Builder {
	b := &Builder{env: env}
	b.splitChunks()
	return b
}

func (b *Builder) splitChunks() {
	f, err := os.Open(b.env.zipPath)
	if err != nil {
		log.Fatalln("Failed to open archive for splitting.", err)
	}
	defer f.Close()
	cs := []*Chunk{}
	ci := 0
	for {
		c := &Chunk{env: b.env, buf: make([]byte, b.env.buildChunkSize), index: ci}
		n, err := f.Read(c.buf)
		if err != nil && err != io.EOF {
			log.Fatalln("Failed to read archive for splitting.", err)
		}
		cs = append(cs, c)
		if n == 0 {
			break
		}
		ci++
	}
	b.chunks = cs
}

func (b *Builder) Generate() {
	ch := make(chan *Chunk)
	for _, c := range b.chunks {
		go c.createData(ch)
	}
	rcs := 0
	for rcs < len(b.chunks) {
		<-ch
		rcs++
	}
}

func (b *Builder) CopyMain() {
	dst, err := os.Create(b.env.buildPkgPath("main.go"))
	if err != nil {
		log.Fatalln("Failed to create main.go.", err)
	}
	src := b.env.mainTemplatePath
	tmpl, err := template.ParseFiles(src)
	if err != nil {
		log.Fatalln("Failed to parse main template.", err)
	}
	if err := tmpl.Execute(dst, b.chunks); err != nil {
		log.Fatalln("Failed to write main template.", err)
	}
}

func (b *Builder) Compile() {
	files, err := filepath.Glob(b.env.buildPkgPath(b.env.buildFuncPrefix + "*.go"))
	if err != nil {
		log.Fatalln("Failed to glob our bindata.", err)
	}
	main := b.env.buildPkgPath("main.go")
	args := append([]string{"build", "-o", b.env.buildBinExec, main}, files...)
	cmd := exec.Command("go", args...)
	cmd.Env = []string{"GOPATH=" + b.env.buildDir}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("Failed run go.", err, string(out))
	}
}
