package main

import (
	"log"
	"path/filepath"
	"os/exec"
)

func Update(env *Environ) error {
	cmd := exec.Command(env.dartPubPath, "update")
	cmd.Dir = env.dartPkgPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err, string(out))
	}
	cmd.Wait()
	return nil
}


func Bundle(env *Environ) error {
	cmd := exec.Command("tar", "cvf",  env.zipPath, "-C", env.dartPkgPath, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(string(out), err)
	}
	cmd = exec.Command("tar", "uvf",  env.zipPath, "-C", filepath.Dir(env.dartVmPath), "dart")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(string(out), err)
	}
	return nil
}
