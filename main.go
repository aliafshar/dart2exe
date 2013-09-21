package main

import (
	"log"
	"path/filepath"
)

func main() {
	env := NewEnviron()
	// Update(env)
	Bundle(env)
	builder := NewBuilder(env)
	builder.Generate()
	builder.CopyMain()
	builder.Compile()
	exec, err := filepath.Rel(env.cwd, env.buildBinExec)
	if err != nil {
		exec = env.buildBinExec
	}
	log.Println("Generated:", exec)
}
